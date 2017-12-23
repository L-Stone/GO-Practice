package controllers

import (
	"github.com/astaxie/beego"
	"github.com/bolgPractice/models"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	op := this.Input().Get("op")
	switch op {
	case "add":
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
	case "del":
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}
	}

	this.TplName = "category.html"
	this.Data["IsCategory"] = true
	var err error
	this.Data["Categories"], err = models.GetAllCategories()

	if err != nil {
		beego.Error(err)
	}
}
