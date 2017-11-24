package models

import (
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func Init() {
	dbHost := beego.AppConfig.String("db:host")
	dbUser := beego.AppConfig.String("db:user")
	dbPwd := beego.AppConfig.String("db:password")
	dbPort := beego.AppConfig.String("db:port")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	if dbPort == "" {
		dbPort = "3306"
	}
	dsn := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbname + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc" + url.QueryEscape(timezone)
	}
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(Auth))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}
