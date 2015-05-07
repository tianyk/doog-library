package models

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"time"
)

var orm *xorm.Engine

func SetEngine() *xorm.Engine {
	var err error
	server := beego.AppConfig.String("mysqlurls")
	username := beego.AppConfig.String("mysqluser")
	password := beego.AppConfig.String("mysqlpass")
	dbName := beego.AppConfig.String("mysqldb")
	orm, err = xorm.NewEngine("mysql", username+":"+password+"@tcp("+server+":3306)/"+dbName+"?charset=utf8")
	if err != nil {
		log.Panicf("SetEngine Error %v", err)
	}
	orm.TZLocation = time.Local
	orm.ShowSQL = false
	return orm
}

func init() {
	SetEngine()
}
