package controllers

import (
	"fmt"
	"myBlog/models"
)

type LoginUserController struct {
	BaseController
}

func (this *LoginUserController) Get() {
	check := this.isLogin
	if check {
		fmt.Println("-------------------LoginUserController--GET--当前用户已经登陆，直接跳转主页")
		this.Redirect("/article", 302)
	} else {
		fmt.Println("-------------------LoginUserController--GET--当前用户没有登陆")
		this.TplName = "login.tpl"
	}
}

func (this *LoginUserController) Post() {
	phone := this.GetString("phone")
	password := this.GetString("password")

	if "" == phone {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写手机号"}
		this.ServeJSON()
	}

	if "" == password {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写密码"}
		this.ServeJSON()
	}

	err, user := models.LoginUser(phone, password)

	if err == nil {
		this.SetSession("userLogin", "1")
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "贺喜你，登录成功", "user": user}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
	}
	this.ServeJSON()
}
