package controllers

import (
	"fmt"
	"mybeetest/lib"
	"mybeetest/models"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type LoginController struct {
	BaseController
}

func (self *LoginController) LoginIn() {
	fmt.Println("-------loginIn into----------")
	if self.userId > 0 {
		self.redirect(beego.URLFor("HomeContrller.Index"))
	}
	beego.ReadFromRequest(&self.Controller)

	if self.isPost() {
		username := strings.TrimSpace(self.GetString("username"))
		password := strings.TrimSpace(self.GetString("password"))
		if username != "" && password != "" {
			user, err := models.AdminGetByName(username)
			fmt.Println("userInfo----george518")
			flash := beego.NewFlash()
			errMsg := ""
			if err != nil || user.Password != lib.MD5([]byte(password+user.Salt)) {
				errMsg = "账号或密码错误"
			} else if user.Status == -1 {
				errMsg = "该账号已禁用"
			} else {
				user.LastIp = self.getClientIp()
				user.LastLogin = time.Now().Unix()
				user.Update()
				authkey := lib.MD5([]byte(self.getClientIp() + "|" + user.Password + user.Salt))
				self.Ctx.SetCookie("authT", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)
				self.redirect(beego.URLFor("HomeController.Index"))
			}
			flash.Error(errMsg)
			flash.Store(&self.Controller)
			self.redirect(beego.URLFor("LoginController.LoginIn"))
		}
	}
	self.TplName = "login/login.html"
	fmt.Println("-------loginIn end----------")
}

//登出
func (self *LoginController) LoginOut() {
	self.Ctx.SetCookie("authT", "")
	self.redirect(beego.URLFor("LoginController.LoginIn"))
}

func (self *LoginController) NoAuth() {
	self.Ctx.WriteString("没有权限")
}
