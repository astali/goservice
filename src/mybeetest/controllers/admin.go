package controllers

type AdminController struct {
	BaseController
}

func (self *AdminController) List() {
	self.Data["pageTitle"] = "管理员管理"
	self.display()
	//self.TplName = "admin/list.html"
}
