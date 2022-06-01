package app

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func Error(c *beego.Controller, code int, err error, msg string) {
	var res Response
	res.Code = code
	res.Message = err.Error()
	if msg != "" {
		res.Message = msg
	}
	logs.Error(res.Message)
	c.Data["json"] = res
	c.ServeJSON(true)
	c.StopRun()
}

func Success(c *beego.Controller, data interface{}, msg string) {
	var res Response
	res.Code = 200
	res.Message = "操作成功"
	res.Data = data
	if msg != "" {
		res.Message = msg
	}
	c.Data["json"] = res
	c.ServeJSON(true)
	c.StopRun()
}
