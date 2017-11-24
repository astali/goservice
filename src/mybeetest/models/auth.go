package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Auth struct {
	Id         int
	Pid        int
	AuthName   string
	AuthUrl    string
	Sort       int
	Icon       string
	IsShow     int
	UserId     int
	CreateId   int
	UpdateId   int
	Status     int
	CreateTime int64
	UpdateTime int64
}

func (a *Auth) TableName() string {
	return TableName("uc_auth")
}

func AuthGetList(page, pageSize int, filters ...interface{}) ([]*Auth, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Auth, 0)
	query := orm.NewOrm().QueryTable(TableName("uc_auth"))
	fmt.Println("---->>query:", query)
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("pid", "sort").Limit(pageSize, offset).All(&list)
	return list, total
}
