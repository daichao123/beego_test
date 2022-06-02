package main

import (
	"beego_test/components"
	"beego_test/models"
	_ "beego_test/routers"
	"beego_test/validate"
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"runtime"
)

func init() {
	//初始化日志
	components.InitLog()
	//链接数据库
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/golang_db?charset=utf8")
	orm.RegisterModel(new(models.Users))

	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true //开启debug 模式
	}
	beego.BConfig.RecoverPanic = true
	beego.BConfig.RecoverFunc = RecoverPanic
	//链接redis

}

func RecoverPanic(c *context.Context, config *beego.Config) {
	if err := recover(); err != nil {
		c.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		//c.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		var stack []string
		for i := 1; ; i++ {
			_, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			//logs.Critical(fmt.Sprintf("%s:%d", file, line))
			stack = append(stack, fmt.Sprintln(fmt.Sprintf("%s:%d", file, line)))
		}
		//显示错误
		data := map[string]interface{}{
			"ret":           4000,
			"AppError":      fmt.Sprintf("%v", err),
			"RequestMethod": c.Input.Method(),
			"RequestURL":    c.Input.URI(),
			"RemoteAddr":    c.Input.IP(),
			"Stack":         stack,
			"GoVersion":     runtime.Version(),
		}
		_ = c.Output.JSON(data, true, true)
		if c.Output.Status != 0 {
			c.ResponseWriter.WriteHeader(c.Output.Status)
		} else {
			c.ResponseWriter.WriteHeader(500)
		}

	}
}

func main() {

	validate.SetDefaultMessage()
	beego.Run()
}
