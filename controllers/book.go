// BookController
package controllers

import (
	. "doog-library/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"log"
	"strconv"
)

type BookController struct {
	beego.Controller
}

// 获取图书信息
func (this *BookController) View() {
	book := new(Book)
	this.ParseForm(book)

	err := book.GetById()
	checkError(err)

	this.Data["json"] = book
	this.ServeJson()
}

// 获取图书的拥有情况
func (this *BookController) Collections() {
	id := this.Ctx.Input.Param(":id")
	bookId, err := strconv.ParseInt(id, 10, 64)
	checkError(err)

	var collections []*Collection
	collections, err = new(Collection).FindCollectionsByBookId(bookId)
	checkError(err)

	this.Data["json"] = collections
	this.ServeJson()
}

// 录入一份新书信息
func (this *BookController) Create() {
	book := new(Book)
	err := this.ParseForm(book)
	checkError(err)

	valid := validation.Validation{}
	var b bool
	b, err = valid.Valid(book)
	checkError(err)

	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		this.Data["json"] = err
		this.ServeJson()
		return
	}
	err = book.Save()
	checkError(err)

	this.Data["json"] = book
	this.ServeJson()
}

// 查看图书列表信息
func (this *BookController) Page() {
	var (
		err error
		// q     string
		// tag   string
		start    int
		count    int
		total    int64
		bookPage []*Book
	)
	start, err = this.GetInt("start")
	checkError(err)
	count, err = this.GetInt("count")
	checkError(err)

	total, bookPage, err = new(Book).Page( /*q, tag, */ start, count)

	page := Page{Start: start, Count: count, Total: total, Data: bookPage}
	this.Data["json"] = &page
	this.ServeJson()
}
