package validate

import (
	"errors"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"reflect"
)

type User struct {
	Username       string `valid:"Required;MaxSize(20);Unique"`
	Password       string `valid:"Required;MaxSize(20);CheckPassword"`
	InviteCode     string `valid:"Required;MaxSize(20)"`
	Mobile         string `valid:"Required;MaxSize(20)"`
	AuthCode       string `valid:"Required;MaxSize(6)"`
	RepeatPassword string `valid:"Required;MaxSize(20)"`
}

func (validateUser *User) ValidateUser() (error error) {
	valid := validation.Validation{}
	b, err := valid.Valid(validateUser)
	if err != nil {
		logs.Error(err, "初始化验证器错误")
	}
	if !b {
		//获取到验证的结构体
		ref := reflect.TypeOf(User{})
		for _, e := range valid.Errors {
			ref.FieldByName(e.Field)
			return errors.New(e.Message)
		}
	}
	return nil
}
