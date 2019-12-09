package routers

import (
	"github.com/astaxie/beego"
	"project/booksys/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
