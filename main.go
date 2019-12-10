package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/siddontang/go/log"
	"project/booksys/common"
	"project/booksys/models/dao"
	_ "project/booksys/models/entity"
	_ "project/booksys/routers"
)

func main() {
	// 初始化数据库连接
	orm.Debug = true
	err := common.DBInit()
	if err != nil {
		log.Error("连接数据库出错, error: ", err)
		return
	}

	// 初始化数据库
	dao.Init()

	//初始化redis服务
	common.RedisInit()

	beego.Run("127.0.0.1:12019")
}
