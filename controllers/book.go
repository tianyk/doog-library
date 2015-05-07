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
	bookId := this.Ctx.Input.Param(":id")
	book := Book{BookId: bookId}
	if err := book.GetById(); err != nil {
		log.Panicf("GetById Error %v", err)
	}
	this.Data["json"] = book
	this.ServeJson()
}

// 获取图书的拥有情况
func (this *BookController) Collections() {
	bookId := this.Ctx.Input.Param(":id")
	collections, err := new(Collection).FindCollectionsByBookId(bookId)
	if err != nil {
		log.Panicf("GetById Error %v", err)
	}
	this.Data["json"] = collections
	this.ServeJson()
}
