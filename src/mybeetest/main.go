package main

import (
	"mybeetest/controllers" //自身业务包
	_ "mybeetest/routers"

	"github.com/astaxie/beego"
	"github.com/beego/admin" //admin 包
)

func init() {
	admin.Run()
	beego.Router("/", &controllers.MainController{})
}

func main() {
	beego.Run()
}
