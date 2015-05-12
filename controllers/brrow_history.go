package controllers

import (
	. "doog-library/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"log"
	// "strconv"
)

type BrrowHistoryController struct {
	beego.Controller
}

func (this *BrrowHistoryController) Create() {
	brrowHistory := new(BrrowHistory)
	err := this.ParseForm(brrowHistory)

	checkError(err)

	valid := validation.Validation{}
	var b bool
	b, err = valid.Valid(brrowHistory)
	checkError(err)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		this.Data["json"] = err
		this.ServeJson()
		return
	}

	err = brrowHistory.Save()
	checkError(err)

	this.Data["json"] = brrowHistory
	this.ServeJson()
}
