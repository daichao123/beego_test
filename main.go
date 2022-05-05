package main

import (
	"beego_test/components"
	"beego_test/models"
	_ "beego_test/routers"
	"beego_test/validate"
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//初始化日志
	components.InitLog()
	//链接数据库
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/golang_db?charset=utf8")
	orm.RegisterModel(new(models.Users))

	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true //开启debug 模式
	}

	//链接redis

}

func main() {

	validate.SetDefaultMessage()
	beego.Run()
}
