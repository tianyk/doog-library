package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	. "github.com/tianyk/doog-library/models"
	"log"
)

type CollectionController struct {
	beego.Controller
}

//
func (this *CollectionController) Create() {
	collection := new(Collection)
	err := this.ParseForm(collection)
	checkError(err)

	valid := validation.Validation{}
	var b bool
	b, err = valid.Valid(collection)
	checkError(err)

	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		this.Data["json"] = err
		this.ServeJson()
		return
	}

	err = collection.Save()
	checkError(err)

	this.Data["json"] = collection
	this.ServeJson()
}
