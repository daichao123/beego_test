package middleware

import (
	"beego_test/tools/app"
	"fmt"
	"github.com/astaxie/beego/context"
	beego "github.com/beego/beego/v2/server/web"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func CustomError(c *beego.Controller) {
	defer func() {
		if err := recover(); err != nil {
			switch errStr := err.(type) {
			case string:
				p := strings.Split(errStr, "#")
				if len(p) == 3 && p[0] == "Customer" {
					statusCode, e := strconv.Atoi(p[1])
					if e != nil {
						break
					}
					fmt.Println(
						time.Now().Format("\n 2006-01-02 15:04:05.9999"),
						"[ERROR]",
						c.Ctx.Request.Method,
						c.Ctx.Request.URL,
						statusCode,
						c.Ctx.Request.RequestURI,
						c.Ctx.Request.RemoteAddr,
						p[2],
					)
					var response app.Response
					response.Code = statusCode
					response.Message = p[2]
					c.Data["json"] = response
					c.ServeJSON(true)
					c.StopRun()
				}
			default:
				panic(err)
			}
		}
	}()

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
