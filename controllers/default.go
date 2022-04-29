package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	mysqluser, err := beego.AppConfig.String("mysqluser1")
	if err != nil {
		logs.Error(err)
	}
	c.Ctx.Output.Body([]byte(mysqluser))
	fmt.Printf("mysqluser:%s\n", mysqluser)
	//fmt.Printf("1111")
	c.Data["Website"] = "beego.test"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
