package models

import (
	"beego_test/validate"
	"github.com/beego/beego/v2/client/orm"
	"github.com/pkg/errors"
)

type Users struct {
	Id       int
	Username string
	Password string
}

// AddUser insertUser
func (*Users) AddUser(users *validate.User) error {
	newOrm := orm.NewOrm()
	newOrm.Begin()
	count, _ := newOrm.QueryTable("users").Filter("invite_code", users.InviteCode).Count()
	if count <= 0 {
		return errors.New("邀请码错误")
	}

	return nil
}
