package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type BaseController struct {
	beego.Controller
	isLogin bool
}

func (this *BaseController) Prepare() {
	userLogin := this.GetSession("userLogin")
	if userLogin == nil {
		this.isLogin = false
	} else {
		this.isLogin = true
	}
	fmt.Println("-------------------BaseController--Prepare")
	this.Data["isLogin"] = this.isLogin
}

func (this *BaseController) Go404() {
	fmt.Println("-------------------BaseController--Go404")
	this.TplName = "404.tpl"
	return
}

func (this *BaseController) GoIndex() {
	fmt.Println("-------------------BaseController--GoIndex")
	this.TplName = "index.tpl"
	return
}

func (this *BaseController) GoArticle() {
	fmt.Println("-------------------BaseController--GoArticle")
	this.TplName = "index.tpl"
	return
}