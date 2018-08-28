package controllers

import "fmt"

//添加blog
type TestController struct {
	BaseController
}

func (this *TestController) TestLogin() {
	fmt.Println("-------------------BaseController--Prepare")
	this.TplName = "login.tpl"
}