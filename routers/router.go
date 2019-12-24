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
			// 注册
			beego.NSRouter("/register", &controllers.AdminControllers{}, "post:Register"),
			// 登陆
			beego.NSRouter("/login", &controllers.AdminControllers{}, "post:Login"),
			// 填写个人信息
			beego.NSRouter("/myinfo", &controllers.AdminControllers{}, "post:MyInfo"),
			// 根据书名查询图书信息
			//beego.NSRouter("/bookname", &controllers.AdminControllers{}, "get:BookName"),
			//// 根据作者查询图书信息
			//beego.NSRouter("/bookauthor", &controllers.AdminControllers{}, "get:BookAuthor"),
			//// 学生借书
			//beego.NSRouter("/borrow", &controllers.AdminControllers{}, "post:Borrow"),
			//// 学生还书
			//beego.NSRouter("/escheat", &controllers.AdminControllers{}, "post:Escheat"),
			//// 查看学生当前借阅信息
			//beego.NSRouter("/borrowinfo", &controllers.AdminControllers{}, "get:BorrowInfo"),
			//// 查看黑名单中的学生
			//beego.NSRouter("/stublacklist", &controllers.AdminControllers{}, "get:StuBlackList"),
			//// 从黑名单删除学生
			//beego.NSRouter("/stublacklistdel", &controllers.AdminControllers{}, "post:StuBlackListDel"),
		),
		//beego.NSNamespace("/student",
		//	// 注册
		//	beego.NSRouter("/register", &controllers.StudentControllers{}, "post:Register"),
		//	// 登陆
		//	beego.NSRouter("/login", &controllers.StudentControllers{}, "post:Login"),
		//	// 填写个人信息
		//	beego.NSRouter("/myinfo", &controllers.StudentControllers{}, "post:MyInfo"),
		//	// 根据书名查询图书信息
		//	beego.NSRouter("/bookname", &controllers.StudentControllers{}, "get:BookName"),
		//	// 根据作者查询图书信息
		//	beego.NSRouter("/bookauthor", &controllers.StudentControllers{}, "get:BookAuthor"),
		//	// 查看当前借阅信息
		//	beego.NSRouter("/borrowinfo", &controllers.StudentControllers{}, "get:BorrowInfo"),
		//	// 续借
		//	beego.NSRouter("/borrowagain", &controllers.StudentControllers{}, "post:BorrowAgain"),
		//),
	)

	beego.AddNamespace(ns)
}

func before(ctx *context.Context) {
	//set output Content-Type to be json
	ctx.Output.Header("Content-Type", "application/json;charset=utf-8")
}
