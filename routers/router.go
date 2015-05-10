package routers

import (
	. "doog-library/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", new(MainController))

	beego.Router("/book", new(BookController), "get:Page")
	beego.Router("/book/:id(^[0-9]+$)", new(BookController), "get:View")
	beego.Router("/book", new(BookController), "post:Create")
	beego.Router("/book/:id(^[0-9]+$)/collections", new(BookController), "get:Collections")

	beego.Router("/collection", new(CollectionController), "post:Create")
}
