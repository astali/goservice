package routers

import (
	"mybeetest/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.LoginController{}, "*:LoginIn")

}
