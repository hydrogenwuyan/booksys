package main

import (
	"github.com/astaxie/beego"
	"github.com/siddontang/go/log"
	"project/booksys/common"
	"project/booksys/models/dao"
	_ "project/booksys/models/entity"
	_ "project/booksys/routers"
	"project/booksys/utils/idutils"
)

func main() {
	// 初始化数据库连接
	err := common.Init()
	if err != nil {
		log.Error("连接数据库出错, error: ", err)
		return
	}

	// 初始化数据库
	dao.Init()

	// 设置id服务
	idutils.SetupWorker(1)

	beego.Run("127.0.0.1:12019")
}

