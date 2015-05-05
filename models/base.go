package model

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	. "doog-library/common"
	"time"
)

var orm *xorm.Engine

func SetEngine() *xorm.Engine {
	beego.Info("db initializing...")
	var err error
	server := beego.AppConfig.String("mysqlurls")
	username := beego.AppConfig.String("mysqluser")
	password := beego.AppConfig.String("mysqlpass")
	dbName := beego.AppConfig.String("mysqldb")
	orm, err = xorm.NewEngine("mysql", username+":"+password+"@tcp("+server+":3306)/"+dbName+"?charset=utf8")
	PanicIf(err)
	orm.TZLocation = time.Local
	orm.ShowSQL = false //beego.AppConfig.String("mysqldb")
	// orm.Logger = xorm.NewSimpleLogger(beego.GetWriter())
	return orm
}

// type DbUtil struct{}

// func (self *DbUtil) GetRecentComments() (comments []Comment) {
// 	err := orm.OrderBy("create_date desc").Limit(5, 0).Find(&comments, &Comment{})
// 	PanicIf(err)
// 	return
// }

// func (self *DbUtil) GetAllLinks() (links []Link) {
// 	err := orm.OrderBy("create_date desc").Find(&links, &Link{})
// 	PanicIf(err)
// 	return links
// }
