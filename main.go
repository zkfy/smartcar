package main

import "github.com/micro-plat/hydra/hydra"
import _ "github.com/go-sql-driver/mysql"

var app = hydra.NewApp(
	hydra.WithPlatName("smartcar"),
	hydra.WithSystemName("smartcar"),
	hydra.WithServerTypes("api"),
	hydra.WithDebug())

//CGO_ENABLED=0 GOOS=linux GOARCH=arm go install
func main() {
	app.Start()
}
