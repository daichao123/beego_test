package controllers

import (
	"beego_test/validate"
	beego "github.com/beego/beego/v2/server/web"
)

type JsonReturn struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
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
		c.Data["json"] = JsonReturn{
			Code:    10003,
			Message: err,
			Data:    nil,
		}
		c.ServeJSON(true)
		c.StopRun()
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
