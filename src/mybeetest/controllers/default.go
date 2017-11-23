package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Astali"] = "astali.com"
	//	c.Ctx.WriteString("----------hello-----")
	c.TplName = "index.tpl"
}