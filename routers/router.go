package routers

import (
	"beego_test/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/reverseList", &controllers.TestController{}, "get:Test")   //反转列表
	beego.Router("/register", &controllers.UserController{}, "post:Register") //注册
	beego.Router("/login", &controllers.UserController{}, "post:Login")       //登录
	beego.Router("/profile", &controllers.UserController{}, "get:GetProfile") //获取个人资料
}
