package models

import (
	"beego_test/validate"
	"github.com/beego/beego/v2/client/orm"
	"github.com/pkg/errors"
	"time"
)

type Users struct {
	Id            int       `json:"id" orm:"id" example:"10"`
	Username      string    `json:"username" orm:"size(60)"`
	Password      string    `json:"password" orm:"size(20)"`
	Encrypt       string    `json:"encrypt" orm:"size(20)"`
	Email         string    `json:"email" orm:"size(30)"`
	Mobile        string    `json:"mobile" orm:"size(20)"`
	RegisterIp    string    `json:"register_ip" orm:"size(20)"`
	IsMainAccount bool      `json:"is_main_account" orm:"size(20)"`
	CreatedAt     time.Time `json:"created_at" orm:"auto_now_add;type(timestamp)"`
	UpdatedAt     time.Time `json:"updated_at" orm:"auto_now_add;type(timestamp)"`
}

// AddUser insertUser
func (*Users) AddUser(users *validate.User) error {
	newOrm := orm.NewOrm()
	//newOrm.Begin()
	count, _ := newOrm.QueryTable("users").Filter("invite_code", users.InviteCode).Count()
	if count <= 0 {
		return errors.New("邀请码错误")
	}

	return nil
}
