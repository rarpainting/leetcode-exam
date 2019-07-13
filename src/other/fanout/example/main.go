package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/rarpainting/glog"

	"mesdata"
	"other/fanout"
)

const (
	minTimeout  = 1000 * time.Millisecond
	idleTimeout = -1 * time.Millisecond
)

var (
	pAddr    = flag.String("p-addr", "192.168.0.1:502", "")
	pSlaveID = flag.Uint("p-slaveid", 1, "")
	lAddr    = flag.String("l-addr", "192.168.0.2:502", "")
	lSlaveID = flag.Uint("l-slaveid", 1, "")

	Addr = flag.String("port", ":11523", "服务端地址")
)

func main() {
	flag.Parse()

	//////////////////////////////////////////////
	pInterrupt := ConnectModbus(p, *pAddr, byte(*pSlaveID), func(controller *fanout.MBController) (chan<- bool, error) {
		return controller.Connect(false, 0, 2, func(b []byte) (res interface{}, err error) {
			if len(b) < 2*2 {
				return nil, mesdata.ErrTooLittle
			}
			p, err := mesdata.TranslatePment(b)
			if err != nil {
				return nil, fmt.Errorf("P Error: %v", err)
			}
			return p, nil
		})
	})
	defer func() {
		pInterrupt <- true
	}()

	//////////////////////////////////////////////
	// 在线/离线 叠片设备
	lInterrupt := ConnectModbus(l, *lAddr, byte(*lSlaveID),
		func(controller *fanout.MBController) (chan<- bool, error) {
			return controller.Connect(false, 0, 2, func(b []byte) (interface{}, error) {
				glog.Debugln(":recieve ", b)
				return mesdata.TranslateL(b)
			})
		})
	defer func() {
		lInterrupt <- true
	}()

	/////////////////////////////////////////////////
	// TCP 连接
	lsInterrupt := ConnectTCP(laser, *lAddr,
		func(controller *fanout.Controller) (chan<- bool, error) {
			return controller.Connect(false,
				func(conn net.Conn, b []byte) (interface{}, int, error) {
					n, err := conn.Read(b)
					if err != nil {
						// 重连
						return nil, 0, err
					}
					laser, err := mesdata.BytesToLaser(b[:n])
					if err != nil {
						return nil, n, err
					}
					return laser, n, err
				})
		})
	defer func() {
		lsInterrupt <- true
	}()

	server := http.Server{
		Addr:         *Addr,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		Handler:      NewRouter(),
	}

	glog.Fatalln(server.ListenAndServe())
}

func ConnectModbus(controller *MBController, mbAddr string, mbSlaveID byte, connectHandle func(*fanout.MBController) (chan<- bool, error)) chan<- bool {
	interrupt := make(chan bool)
	go func() {
		fanoutController, err := fanout.NewModbusController(mbAddr, mbSlaveID, minTimeout, idleTimeout)
		if err != nil {
			glog.Fatalln("::", mbAddr, ":new controller err: ", err.Error())
		}
		glog.Infoln(mbAddr, ":new controller ok")

		sinterrupt, err := make(chan<- bool), nil
		for {
			sinterrupt, err = connectHandle(fanoutController)
			if err != nil {
				time.Sleep(50 * time.Millisecond)
				continue
			}
			break
		}

		controller.C = fanoutController
		for {
			select {
			case <-interrupt:
				close(interrupt)
				sinterrupt <- true
				fanoutController.DisConnect()
				return
			default:
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	return interrupt
}

func ConnectTCP(controller *Controller, addr string, connectHandle func(*fanout.Controller) (chan<- bool, error)) chan<- bool {
	interrupt := make(chan bool)
	go func() {
		fanoutController, err := fanout.NewController(addr, minTimeout)
		if err != nil {
			glog.Fatalln("::", addr, ":new controller err: ", err.Error())
		}
		glog.Infoln(addr, ":new controller ok")

		sinterrupt, err := make(chan<- bool), nil
		for {
			sinterrupt, err = connectHandle(fanoutController)
			if err != nil {
				time.Sleep(50 * time.Millisecond)
				continue
			}
			break
		}

		controller.C = fanoutController
		for {
			select {
			case <-interrupt:
				close(interrupt)
				sinterrupt <- true
				fanoutController.DisConnect()
				return
			default:
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	return interrupt
}
