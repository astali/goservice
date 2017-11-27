package controllers

type AuthController struct {
	BaseController
}

func (self *AuthController) Index() {
	self.Data["pageTitle"] = "权限因子"
	self.display()
}
