package fanout

import (
	"fmt"
	"sync"
	"time"

	"github.com/goburrow/modbus"
	"github.com/rarpainting/glog"
	gouuid "github.com/satori/go.uuid"
)

// MBSession 会话
type MBSession struct {
	c modbus.Client
	h *modbus.TCPClientHandler
}

// NewModbusSession 创建会话
func NewModbusSession(address string, slaveid byte, timeout, idleTimeout time.Duration) (*MBSession, error) {
	mintime := MinTimeout / time.Millisecond
	maxtime := MaxTimeout / time.Millisecond

	if timeout < MinTimeout || idleTimeout > MaxTimeout {
		return nil, fmt.Errorf("timeout must between on %dms to %dms", mintime, maxtime)
	} else if idleTimeout > 0 && (idleTimeout < MinTimeout || idleTimeout > MaxTimeout) {
		return nil, fmt.Errorf("idleTimeout must between on %dms to %dms OR less than 0", mintime, maxtime)
	}

	handler := modbus.NewTCPClientHandler(address)

	// timeout
	// 至少要 500ms , 但是连接后容易获取不到数据
	handler.Timeout = timeout // 100 * time.Millisecond

	// idletimeout
	handler.IdleTimeout = idleTimeout

	handler.SlaveId = slaveid

	c := modbus.NewClient(handler)
	return &MBSession{
		c: c,
		h: handler,
	}, nil
}

// Connect 单单连接上, 而没有检测
func (s *MBSession) Connect() error {
	return s.h.Connect()
}

// ConnectBlock 阻塞 启动通道
func (s *MBSession) ConnectBlock() {
	for {
		if err := s.Connect(); err != nil {
			time.Sleep(time.Second * 1)
			continue
		} else {
			return
		}
	}
}

func (s *MBSession) Close() error {
	return s.h.Close()
}

// 把常用的复制出去

func (s *MBSession) ReadCoils(addr, quantity uint16) (result []byte, err error) {
	return s.c.ReadCoils(addr, quantity)
}

func (s *MBSession) ReadDiscreteInputs(addr, quantity uint16) (result []byte, err error) {
	return s.c.ReadDiscreteInputs(addr, quantity)
}

func (s *MBSession) ReadHoldingRegisters(addr, quantity uint16) (result []byte, err error) {
	return s.c.ReadHoldingRegisters(addr, quantity)
}

type read struct {
	Disc []byte
	Coil []byte
	Rhr  []byte
}

// type write struct {
// 	Ch           chan interface{}
// 	ShouldCancel bool
// }

// Controller 通道结构体
type MBController struct {
	s *MBSession

	read chan interface{} // 创建的 输入 通道列表

	write map[string]*write // 缓存下来的 输出 通道列表

	mutex sync.Mutex // 对 []write 的锁
}

// NewModbusController 设置 Modbus 连接会话
func NewModbusController(address string, slaveid byte, timeout, idleTimeout time.Duration) (*MBController, error) {
	s, err := NewModbusSession(address, slaveid, timeout, idleTimeout)
	if err != nil {
		return nil, err
	}
	return &MBController{
		s:     s,
		write: make(map[string]*write), // write map
		mutex: sync.Mutex{},            // lock
	}, nil
}

// Connect 开启实际连接
func (c *MBController) Connect(async bool, address, quantity uint16, handleFunc func([]byte) (interface{}, error)) (
	interrupt chan<- bool, err error) {
	c.read = make(chan interface{})
	s := c.s
	if err := s.Connect(); err != nil {
		return nil, err
	}

	interrupt = c.fanOut(c.read, async) // Tee 协程 // NOTE: 关闭 read channel

	go func() {
		defer s.Close()
		defer c.CancelRead()
		defer func() { // 防止进程中断
			if err := recover(); err != nil {
				// 未知错误导致 panic
				// TODO: 回报到调用者
				glog.Error("![Unknown Error]! ", err)
			}
		}()

	READ_HOLDING:
		for {
			if len(c.write) == 0 {
				// 没有 输出通道 的时候进入 间隔为 100ms 的休眠
				time.Sleep(100 * time.Millisecond)
				continue READ_HOLDING
			}

			b, err := s.c.ReadHoldingRegisters(address, quantity) // 4X 数据寄存器

			if err != nil {
				glog.Error("ReadHoldingRegisters Error", err.Error())

				s.h.Close()

				// 中途出现连接故障, 则
				// 以 200 << [0~4] ms 的渐慢重连
				for reconnectCount := uint(0); ; {
					time.Sleep(50 * time.Millisecond)

					err = s.h.Connect()
					if err == nil {
						continue READ_HOLDING
					}

					time.Sleep((300 << reconnectCount) * time.Millisecond)

					if reconnectCount < 4 {
						reconnectCount++
					}
				}
			}

			glog.Infof("::connect: mb data is: (want=%d, len=%d) %x\n", quantity, len(b), b)

			// log.Println("::connect: handle data to structure")
			res, err := handleFunc(b)
			if err != nil {
				glog.Errorln("handleFunc Error", err.Error())
				continue READ_HOLDING
			}
			// log.Println("::connect: data -> c.read")
			c.read <- res
			time.Sleep(UpdateRate)
		}
	}()
	return interrupt, nil
}

// DisConnect 关闭实际连接
func (c *MBController) DisConnect() error {
	c.CancelRead() // 关闭读通道
	return c.s.Close()
}

// AddWrite 添加输出通道
func (c *MBController) AddWrite() (string, <-chan interface{}, error) {
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
func (c *MBController) CancelWrite(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if write, ok := c.write[key]; ok {
		write.ShouldCancel = true
		// c.write[key] = write // 强行多刷了一次 write, 是多余的举动吗 ?
	}
}

// CancelRead 关闭读通道
func (c *MBController) CancelRead() {
	defer func() {
		if err := recover(); err != nil {
			glog.Error(err)
		}
	}()
	close(c.read)
}

// Tee 扇出 -- 单输入多输出
// 参数 read 是避免在读取的同时出现写入到 s.read 中的情况
// TODO: 超时关闭
func (c *MBController) fanOut(chRead <-chan interface{}, async bool) chan<- bool {
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
				glog.Error("[Panic]", err)
			}
		}()

		// 从 read 写入到 write
		toWrite := func(read interface{}, write *write) {
			// BUG: 怎么导致出现 空的情况 ?
			if read == nil || write == nil {
				glog.Error("[Error]", "read OR write is nil", "[Read]:", read, "[Write]:", write)
				return
			}

			glog.Infoln("::connect: data->write channel")
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

					glog.Infoln("::connect: data read to write")
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
