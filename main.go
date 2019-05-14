package main

import "github.com/micro-plat/hydra/hydra"
import _ "github.com/go-sql-driver/mysql"

type smartcar struct {
	*hydra.MicroApp
}

//CGO_ENABLED=0 GOOS=linux GOARCH=arm go install
func main() {
	app := &smartcar{
		hydra.NewApp(
			hydra.WithPlatName("smartcar"),
			hydra.WithSystemName("smartcar"),
			hydra.WithServerTypes("api"),
			hydra.WithDebug()),
	}

	app.init()
	app.install()
	app.handling()

	app.Start()
}
