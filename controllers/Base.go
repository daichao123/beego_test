package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

type ReturnMsg struct {
	Code    int                    `json:"code" example:"200"`
	Message string                 `json:"message" example:"success"`
	Data    map[string]interface{} `json:"data"`
}

func (c *BaseController) Json(obj interface{}) {
	c.Data["json"] = obj
	c.ServeJSON(true)
	c.StopRun()
}
