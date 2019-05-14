package services

import (
	"time"

	"github.com/stianeikeland/go-rpio"
)

var (
	fl0 = rpio.Pin(17) //11
	fl1 = rpio.Pin(18) //12
	flA = rpio.Pin(23) //16
)

//Open 打开所有pin
func Open() error {
	if err := rpio.Open(); err != nil {
		return err
	}

	fl0.Output()
	fl1.Output()
	return nil
}

//Start 启动
func Start(t time.Duration) {
	flA.High()
	fl0.High()
	fl1.Low()
	time.Sleep(t)
	flA.Low()
}

//Close 关闭
func Close() {
	rpio.Close()
}
