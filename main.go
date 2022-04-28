package main

import (
	_ "beego_test/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

