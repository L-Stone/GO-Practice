package routers

import (
	"github.com/astaxie/beego"
	"github.com/bolgPractice/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
}
