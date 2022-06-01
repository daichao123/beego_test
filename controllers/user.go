package controllers

import (
	"beego_test/models"
	"beego_test/tools/app"
	"beego_test/validate"
	beego "github.com/beego/beego/v2/server/web"
	"strings"
)

type JsonReturn struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"` //Data字段需要设置为interface类型以便接收任意数据
	//json标签意义是定义此结构体解析为json或序列化输出json时value字段对应的key值,如不想此字段被解析可将标签设为`json:"-"`
}

type UserController struct {
	beego.Controller
}

// Register 用户注册
func (c *UserController) Register() {
	user := validate.User{
		Username:       c.GetString("username"),
		Password:       c.GetString("password"),
		InviteCode:     c.GetString("invite_code"),
		Mobile:         c.GetString("mobile"),
		AuthCode:       c.GetString("auth_code"),
		RepeatPassword: c.GetString("repeat_password"),
	}
	err := user.ValidateUser()
	if err != nil {
		app.Error(&beego.Controller{}, 10003, err, err.Error())
	}
	addr := strings.Split(c.Ctx.Request.RemoteAddr, ":")
	lock := models.Lock(addr[0], 5)
	if lock != 1 {
		app.Error(&beego.Controller{}, 10003, nil, "操作频繁、请稍后再试")
	}
	res := models.CheckVCode(user.Username, user.InviteCode, "register")
	if !res {
		app.Error(&beego.Controller{}, 10003, nil, "邀请码错误")
	}
	//username := c.GetString("username")
	//password := c.GetString("password")
	//valid := validation.Validation{
	//	RequiredFirst: false,
	//	Errors:        nil,
	//	ErrorsMap:     nil,
	//}
	//var message = map[string]string{
	//	"Required": "参数不能为空",
	//}
	//validation.SetDefaultMessage(message)

	//valid.Required()
}
