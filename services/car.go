package services

import (
	"time"

	"github.com/stianeikeland/go-rpio"
)

var (
	fl0 = rpio.Pin(17) //11
	fl1 = rpio.Pin(18) //12
	flA = rpio.Pin(27) //13

	fr0 = rpio.Pin(23) //16
	fr1 = rpio.Pin(24) //18
	frA = rpio.Pin(25) //22


	bl0 = rpio.Pin(5) //29
	bl1 = rpio.Pin(6) //31
	blA = rpio.Pin(12) //32

	br0 = rpio.Pin(13) //33
	br1 = rpio.Pin(19) //35
	brA = rpio.Pin(26) //37

	pins=[]rpio.Pin{
		fl0,fl1,fr0,fr1,bl0,bl1,br0,br1,
	}

)

//Open 打开所有pin
func Open() error {
	if err := rpio.Open(); err != nil {
		return err
	}
	for _,p:=range pins{
		p.Output()
	}
	return nil
}

func flForward(){
	flA.High()
	fl0.High()
	fl1.Low()
}

func flBack(){
	flA.High()
	fl0.Low()
	fl1.High()
}
func flStop(){
	flA.Low()
}

func frForward(){
	frA.High()
	fr0.High()
	fr1.Low()
}
func frBack(){
	frA.High()
	fr0.Low()
	fr1.High()
}
func frStop(){
	frA.Low()
}

func blForward(){
	blA.High()
	bl0.High()
	bl1.Low()
}
func blBack(){
	blA.High()
	bl0.Low()
	bl1.High()
}
func blStop(){
	blA.Low()
}
func brForward(){
	brA.High()
	br0.High()
	br1.Low()
}
func brBack(){
	brA.High()
	br0.Low()
	br1.High()
}
func brStop(){
	brA.Low()
}


//Forward 启动服务
func Forward(t time.Duration) {
	flForward()
	frForward()
	blForward()
	brForward()
	time.Sleep(t)

}
//FLeft 前进左转
func FLeft(t time.Duration){
	flStop()
	frForward()
	blForward()
	brForward()
	time.Sleep(t)
}

//FRight 前进右转
func FRight(t time.Duration){
	flForward()
	frStop()
	blForward()
	brForward()
	time.Sleep(t)
}
//Back 后退
func Back(t time.Duration){
	flBack()
	frBack()
	blBack()
	brBack()
	time.Sleep(t)
}

//BLeft 后退左转
func BLeft(t time.Duration){
	flBack()
	frBack()
	blStop()
	brBack()
	time.Sleep(t)
}

//BRight 后退右转
func BRight(t time.Duration){
	flBack()
	frBack()
	brStop()
	blBack()
	time.Sleep(t)
}




//Stop  停止服务
func Stop(){
	flStop()
	frStop()
	blStop()
	brStop()
}




//Close 关闭
func Close() {
	rpio.Close()
}
