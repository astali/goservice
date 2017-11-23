package routers

import (
	"fmt"
	"mybeetest/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Get("/login/:ids([0-7]+).html", func(ctx *context.Context) {
		fmt.Println("-----", ctx.Input.Param(":ext"))
		ctx.Output.Body([]byte("hello world"))
	})

}
