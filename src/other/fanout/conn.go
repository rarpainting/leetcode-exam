package fanout

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/rarpainting/glog"
	gouuid "github.com/satori/go.uuid"
)

const (
	MinHold = 25
	// MinDisc ReadDiscreteInputs 的最小读取数
	MinDisc = 128
	// MinCoil ReadCoils 的最小读取数
	MinCoil = 128

	MaxTimeout = 3 * time.Second
	MinTimeout = 1000 * time.Millisecond

	UpdateRate      = 500 * time.Millisecond
	KeepAlivePeriod = 100 * time.Millisecond
)

// Session 会话
type Session struct {
	address string
	timeout time.Duration
	conn    *net.TCPConn
}

// NewSession 创建会话
func NewSession(address string, timeout time.Duration) (*Session, error) {
	if timeout < MinTimeout || timeout > MaxTimeout {
		return nil, fmt.Errorf("timeout must between on %v to %v", MinTimeout, MaxTimeout)
	}

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return nil, err
	}
	tcpconn, ok := conn.(*net.TCPConn)
	if !ok {
		return nil, fmt.Errorf("conn is not *TCPConn")
	}
	tcpconn.SetKeepAlive(true)
	tcpconn.SetKeepAlivePeriod(KeepAlivePeriod)

	return &Session{
		address: address,
		timeout: timeout,
		conn:    tcpconn,
	}, nil
}

func (s *Session) Connect() error {
	conn, err := net.DialTimeout("tcp", s.address, s.timeout)
	if err != nil {
		return err
	}
	tcpconn, ok := conn.(*net.TCPConn)
	if !ok {
		return fmt.Errorf("conn is not *TCPConn")
	}

	tcpconn.SetKeepAlive(true)
	tcpconn.SetKeepAlivePeriod(KeepAlivePeriod)
	s.conn = tcpconn
	return nil
}

func (s *Session) Close() error {
	return s.conn.Close()
}

// 把常用的复制出去
func (s *Session) Read(b []byte) (n int, err error) {
	return s.conn.Read(b)
}

func (s *Session) SetDeadline(t time.Time) error {
	return s.conn.SetDeadline(t)
}

func (s *Session) SetReadDeadline(t time.Time) error {
	return s.conn.SetReadDeadline(t)
}

func (s *Session) SetWriteDeadline(t time.Time) error {
	return s.conn.SetWriteDeadline(t)
}

// type read struct {
// 	Disc []byte
// 	Coil []byte
// 	Rhr  []byte
// }

type write struct {
	Ch           chan interface{}
	ShouldCancel bool
}

// Controller 通道结构体
type Controller struct {
	s *Session

	read chan interface{} // 创建的 输入 通道列表

	write map[string]*write // 缓存下来的 输出 通道列表

	mutex sync.Mutex // 对 []write 的锁
}

// NewController 设置 Modbus 连接会话
func NewController(address string, timeout time.Duration) (*Controller, error) {
	s, err := NewSession(address, timeout)
	if err != nil {
		return nil, err
	}
	return &Controller{
		s:     s,
		write: make(map[string]*write), // write map
		mutex: sync.Mutex{},            // lock
	}, nil
}

// Connect 开启实际连接
func (c *Controller) Connect(async bool, readFunc func(net.Conn, []byte) (interface{}, int, error)) (
	interrupt chan<- bool, err error) {
	c.read = make(chan interface{})
	s := c.s

	interrupt = c.fanOut(c.read, async) // Tee 协程 // NOTE: 关闭 read channel

	go func() {
		defer s.Close()
		defer c.CancelRead()

	READ_HOLDING:
		for {
			if len(c.write) == 0 {
				// 没有 输出通道 的时候进入 间隔为 100ms 的休眠
				time.Sleep(100 * time.Millisecond)
				continue READ_HOLDING
			}

			b := make([]byte, 1024)
			iface, n, err := readFunc(c.s.conn, b)

			// 如果出现 error code , 会出现以 50ms 重连的怪像
			if err != nil {
				s.conn.Close()

				// 中途出现连接故障, 则
				// 以 200 << [0~4] ms 的渐慢重连
				for reconnectCount := uint(0); ; {
					time.Sleep(300 * time.Millisecond)

					err = s.Connect()
					if err == nil {
						continue READ_HOLDING
					}

					time.Sleep((300 << reconnectCount) * time.Millisecond)

					if reconnectCount < 4 {
						reconnectCount++
					}
				}
			}

			glog.Debugf("::connect: data is: (len=%d) %x\n", n, b[:n])

			c.read <- iface
			time.Sleep(UpdateRate)
		}
	}()
	return interrupt, nil
}

// DisConnect 关闭实际连接
func (c *Controller) DisConnect() error {
	c.CancelRead() // 关闭读通道
	return c.s.Close()
}

// AddWrite 添加输出通道
func (c *Controller) AddWrite() (string, <-chan interface{}, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// 生成随机字符串
	uuid, err := gouuid.NewV4()
	if err != nil {
		return "", nil, err
	}

	key := uuid.String()
	ch := make(chan interface{})

	c.write[key] = &write{
		Ch: ch,
	}
	return key, ch, nil
}

// CancelWrite 关闭索引 i 的写通道, 索引 i 作为 AddWrite() 的返回值返回
// 如果与 CNC 的读写连接已经断开了, 就没必要再设置了
func (c *Controller) CancelWrite(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if write, ok := c.write[key]; ok {
		write.ShouldCancel = true
		// c.write[key] = write // 强行多刷了一次 write, 是多余的举动吗 ?
	}
}

// CancelRead 关闭读通道
func (c *Controller) CancelRead() {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorln(err)
		}
	}()
	close(c.read)
}

// Tee 扇出 -- 单输入多输出
// 参数 read 是避免在读取的同时出现写入到 s.read 中的情况
// TODO: 超时关闭
func (c *Controller) fanOut(chRead <-chan interface{}, async bool) chan<- bool {
	interrupt := make(chan bool)

	go func() {
		defer func() {
			// 输入关闭, 输出全部关闭
			for k, v := range c.write {
				close(v.Ch)
				delete(c.write, k)
			}
		}()
		defer func() { // 防止进程中断
			if err := recover(); err != nil {
				// 未知错误导致 panic
				// TODO: 回报到调用者
				glog.Errorln("[Panic]", err)
			}
		}()

		// 从 read 写入到 write
		toWrite := func(read interface{}, write *write) {
			// BUG: 怎么导致出现 空的情况 ?
			if read == nil || write == nil {
				glog.Errorln("[Error]", "read OR write is nil", "[Read]:", read, "[Write]:", write)
				return
			}

			glog.Debugln("::connect: data->write channel")
			write.Ch <- read
		}

		for {
			select {
			case <-interrupt:
				return

			case v, ok := <-chRead:
				if !ok { // 通道被关闭
					return
				}
				c.mutex.Lock() // 放在前面部分原因是防止中途 c.write 被改写
				for key, write := range c.write {
					write := write // 意图在哪 ?

					if write.ShouldCancel {
						// 调用 CancelWrite, 在这里真正关闭 通道, 且删除
						close(write.Ch)
						delete(c.write, key)
						continue
					}

					glog.Debugln("::connect: data read to write")
					if async {
						go toWrite(v, write)
					} else {
						toWrite(v, write)
					}
				}
				c.mutex.Unlock()
			default:
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()
	return interrupt
}
