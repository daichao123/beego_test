package controllers

import (
	"beego_test/models"
	"beego_test/tools"
	"beego_test/validate"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
)

type JsonReturn struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"` //Data字段需要设置为interface类型以便接收任意数据
	//json标签意义是定义此结构体解析为json或序列化输出json时value字段对应的key值,如不想此字段被解析可将标签设为`json:"-"`
}

type UserController struct {
	BaseController
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
		fmt.Println(err.Error())
		msg := ReturnMsg{
			Code:    10003,
			Message: strings.Join(strings.Fields(err.Error()), ""),
		}
		c.Json(msg)
	}
	addr := strings.Split(c.Ctx.Request.RemoteAddr, ":")
	lock := models.Lock(addr[0], 5)
	if lock != 1 {
		c.Json(ReturnMsg{
			Code:    10003,
			Message: "操作频繁",
		})
	}
	res := models.CheckVCode(user.Username, user.InviteCode, "register")
	if !res {
		c.Json(ReturnMsg{
			Code:    10003,
			Message: "邀请码错误",
		})
	}
	isMainAccount, _ := c.GetBool("is_main_account")
	isMainAccountInt := 0
	if isMainAccount {
		isMainAccountInt = 1
	}
	users := models.Users{
		Username:      user.Username,
		Password:      tools.GetEncryptStringByMd5(user.Password, tools.GetRandString(8)),
		Encrypt:       "",
		Email:         "",
		Mobile:        user.Mobile,
		RegisterIp:    addr[0],
		IsMainAccount: isMainAccountInt,
	}
	newOrm := orm.NewOrm()
	newOrm.Using("user_service")
	insert, err := newOrm.Insert(&users)
	if err != nil && insert == 0 {
		panic(err)
	}

	c.Json(ReturnMsg{
		Code:    200,
		Message: "操作成功",
		Data:    map[string]interface{}{"id": insert},
	})
}

func (c *UserController) Login() {
	username := c.GetString("username", "")
	password := c.GetString("password", "")
	if username == "" && password == "" {
		c.Json(ReturnMsg{
			Code:    10003,
			Message: "参数错误",
		})
	}
	usersModel := models.Users{
		Username: username,
	}
	newOrm := orm.NewOrm()
	newOrm.Using("user_service")
	err := newOrm.Read(&usersModel)
	if err == orm.ErrNoRows {
		c.Json(ReturnMsg{
			Code:    10003,
			Message: "数据未找到",
		})
	} else if err == orm.ErrMissPK {
		c.Json(ReturnMsg{
			Code:    10003,
			Message: "找不到主键",
		})
	}
	if password != tools.GetEncryptStringByMd5(usersModel.Password, usersModel.Encrypt) {
		c.Json(ReturnMsg{
			Code:    10003,
			Message: "密码错误",
		})
	}
	//使用JWT
}
