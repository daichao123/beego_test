package validate

import (
	"errors"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"reflect"
)

type User struct {
	Username       string `json:"用户名" alias:"用户名" valid:"Required;MaxSize(20);Unique" `
	Password       string `alias:"密码" valid:"Required;MaxSize(20);CheckPassword"`
	InviteCode     string `alias:"邀请码" valid:"Required;MaxSize(20)"`
	Mobile         string `alias:"手机号" valid:"Required;MaxSize(20)" `
	AuthCode       string `alias:"验证码" valid:"Required;MaxSize(6)"`
	RepeatPassword string `alias:"重复密码" valid:"Required;MaxSize(20)"`
}

func (validateUser *User) ValidateUser() (error error) {
	valid := validation.Validation{}
	b, err := valid.Valid(validateUser)
	if err != nil {
		logs.Error(err, "初始化验证器错误")
	}
	if !b {
		//获取到验证的结构体
		//ref := reflect.TypeOf(User{})
		ref := reflect.TypeOf(User{})
		for _, e := range valid.Errors {
			filed, _ := ref.FieldByName(e.Field)
			var alias = filed.Tag.Get("alias")
			return errors.New(alias + e.Message)
		}
	}
	return nil
}
