package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/smartcar/services"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func init() {
	app.Closing(func(c component.IContainer) error {
		services.Close()
		return nil
	})

	app.Initializing(func(c component.IContainer) error {
		if err := services.Open(); err != nil {
			return err
		}
		app.Micro("/car", services.NewCarctlHandler)

		return nil
	})
}
