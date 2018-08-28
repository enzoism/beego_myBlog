package routers

import (
	"myBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.LoginUserController{})
	beego.Router("/404.html", &controllers.BaseController{}, "*:Go404")
	beego.Router("/index.html", &controllers.BaseController{}, "*:GoIndex")
	beego.Router("/article", &controllers.BaseController{}, "*:GoIndex")
	beego.Router("/testLogin", &controllers.TestController{},"*:TestLogin")


}
