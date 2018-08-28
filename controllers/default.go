package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}


func (this *MainController) Go404() {
	fmt.Println("-------------------MainController--Go404")
	this.TplName = "404.tpl"
	return
}