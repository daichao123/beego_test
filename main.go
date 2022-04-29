package main

import (
	"beego_test/components"
	_ "beego_test/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	components.InitLog()
	//logs.SetLogger(logs.AdapterFile, `{"filename":"logs/test.log","level":7,"maxlines":100000,"daily":true,"maxdays":10}`)
}

func main() {
	beego.Run()
}
