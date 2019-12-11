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
			// 登陆
			beego.NSRouter("/login", &controllers.AdminControllers{}, "post:Login"),
			// 填写个人信息
			beego.NSRouter("/myinfo", &controllers.AdminControllers{}, "post:MyInfo"),
			// 根据书名查询图书信息
			beego.NSRouter("/bookname", &controllers.AdminControllers{}, "post:BookName"),
			// 根据作者查询图书信息
			beego.NSRouter("/bookauthor", &controllers.AdminControllers{}, "post:BookAuthor"),
			// 学生借书
			beego.NSRouter("/borrow", &controllers.AdminControllers{}, "post:Borrow"),
			// 学生还书
			beego.NSRouter("/escheat", &controllers.AdminControllers{}, "post:Escheat"),
			// 查看学生当前借阅信息
			beego.NSRouter("/borrowinfo", &controllers.AdminControllers{}, "post:BorrowInfo"),
			// 查看黑名单中的学生
			beego.NSRouter("/", &controllers.AdminControllers{}, "post:BlackList"),
			// 从黑名单删除学生
			beego.NSRouter("/", &controllers.AdminControllers{}, "post:DelBlackList"),
		),
		beego.NSNamespace("/student",
			// 登陆
			beego.NSRouter("/login", &controllers.StudentControllers{}, "post:Login"),
			// 填写个人信息
			beego.NSRouter("/myinfo", &controllers.StudentControllers{}, "post:MyInfo"),
			// 根据书名查询图书信息
			beego.NSRouter("/bookname", &controllers.StudentControllers{}, "post:BookName"),
			// 根据作者查询图书信息
			beego.NSRouter("/bookauthor", &controllers.StudentControllers{}, "post:BookAuthor"),
			// 查看当前借阅信息
			beego.NSRouter("/borrowinfo", &controllers.StudentControllers{}, "post:BorrowInfo"),
			// 续借
			beego.NSRouter("/borrowagain", &controllers.StudentControllers{}, "post:BorrowAgain"),
		),
	)

	beego.AddNamespace(ns)
}

func before(ctx *context.Context) {
	//set output Content-Type to be json
	ctx.Output.Header("Content-Type", "application/json;charset=utf-8")
}
