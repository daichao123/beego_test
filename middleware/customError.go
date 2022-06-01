package middleware

import beego "github.com/beego/beego/v2/server/web"

func CustomError(c *beego.Controller) {
	defer func() {
		if err := recover(); err != nil {

		}
	}()
}
