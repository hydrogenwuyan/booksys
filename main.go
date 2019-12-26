package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	"project/booksys/common"
	"project/booksys/models/dao"
	_ "project/booksys/models/entity"
	_ "project/booksys/routers"
	_ "project/booksys/task"
)

func main() {
	// 初始化日志
	common.LogInit()

	// 初始化数据库连接
	orm.Debug = true
	err := common.DBInit()
	if err != nil {
		common.LogFuncError("连接数据库出错, error: %v", err)
		return
	}

	// 初始化数据库
	dao.Init()

	//初始化redis服务
	common.RedisInit()

	// 执行定时任务
	toolbox.StartTask()
	defer toolbox.StopTask()

	// 开启http服务
	startHttpServer()
}

func startHttpServer() {
	host := "127.0.0.1"
	port := beego.AppConfig.String("httpport")
	addr := fmt.Sprintf("%s:%s", host, port)
	beego.Run(addr)
}
