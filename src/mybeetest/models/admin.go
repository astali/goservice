package models

import (
	"github.com/astaxie/beego/orm"
)

type Admin struct {
	Id         int
	LoginName  string
	RealName   string
	Password   string
	RoleIds    string
	Phone      string
	Email      string
	Salt       string
	LastLogin  int64
	LastIp     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (a *Admin) TableName() string {
	return TableName("uc_admin")
}

func AdminAdd(a *Admin) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func AdminGetById(id int) (*Admin, error) {
	r := new(Admin)
	err := orm.NewOrm().QueryTable(TableName("uc_admin")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func AdminGetByName(loginName string) (*Admin, error) {
	a := new(Admin)
	err := orm.NewOrm().QueryTable(TableName("uc_admin")).Filter("login_name", loginName).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *Admin) Update(filters ...string) error {
	if _, err := orm.NewOrm().Update(a, filters...); err != nil {
		return err
	}
	return nil
}
