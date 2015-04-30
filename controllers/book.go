package controllers

import (
	. "doog-library/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"log"
	// "strconv"
)

type BookController struct {
	beego.Controller
}

/**
 * 获取图书信息
 * @param  {[type]} this *BookController) View( [description]
 * @return {[type]}      [description]
 */
func (this *BookController) View() {
	id := this.Ctx.Input.Param(":id")
	valid := validation.Validation{}
	valid.MinSize(id, 6, "id")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	// id2, _ := strconv.ParseInt(id, 10, 32)
	log.Println(id)
	this.Ctx.WriteString(id)
}

/**
 * 获取图书的拥有情况
 * @param  {[type]} this *BookController) Collections( [description]
 * @return {[type]}      [description]
 */
func (this *BookController) Collections() {
	id := this.Ctx.Input.Param(":id")
	// select * from colections where book_id = id;
	colections := []Collection{Collection{id, "12306", "loan"}, Collection{id, "12306", "lose"}}
	log.Println(id)
	this.Data["json"] = &colections
	this.ServeJson()
}
