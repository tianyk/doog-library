package routers

import (
	"doog-library/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/book/:id(^[0-9]+$)", &controllers.BookController{}, "get:View")
	beego.Router("/book/:id(^[0-9]+$)/collections", &controllers.BookController{}, "get:Collections")
}
