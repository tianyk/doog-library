// BookController
package controllers

import (
	. "doog-library/models"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/validation"
	"log"
	// "strconv"
)

type BookController struct {
	beego.Controller
}

// 获取图书信息
func (this *BookController) View() {
	id := this.Ctx.Input.Param(":id")

	book := NewBook().FindById(id)
	this.Data["json"] = book
	this.ServeJson()
}

// 获取图书的拥有情况
func (this *BookController) Collections() {
	id := this.Ctx.Input.Param(":id")
	log.Printf("Book ID is %s", id)
	collections := NewCollection().FindBookCollectionsByBookId(id)
	this.Data["json"] = collections
	this.ServeJson()
}
