package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"project/booksys/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSBefore(before),
		beego.NSNamespace("/admin",
			beego.NSRouter("/login", &controllers.AdminControllers{}, "post:Login"),
		),
	)

	beego.AddNamespace(ns)
}

func before(ctx *context.Context) {
	//set output Content-Type to be json
	ctx.Output.Header("Content-Type", "application/json;charset=utf-8")
}
