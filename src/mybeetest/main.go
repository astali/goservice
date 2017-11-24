package main

import (
	"mybeetest/models"
	_ "mybeetest/routers"

	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	beego.Run()
}
