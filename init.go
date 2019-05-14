package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/smartcar/services"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func (r *smartcar) init() {
	r.Closing(func(c component.IContainer)error{
		services.Close()
		return nil
	})
	
	r.Initializing(func(c component.IContainer) error {
		if err:=services.Open();err!=nil{
			return err
		}

		//appconf.func#//
		//#appconf.func//

		//db.init#//
		//#db.init//

		//cache.init#//
		//#cache.init//

		//queue.init#//
		//#queue.init//

		//login.router#//
		//#login.router//

		//service.router#//
		r.Micro("/car",services.NewCarctlHandler)
		//#service.router//

		return nil
	})
}
