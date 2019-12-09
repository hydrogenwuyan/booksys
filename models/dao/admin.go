package dao

import (
	"github.com/astaxie/beego/orm"
	"project/booksys/common"
)

type AdminDao struct {
	orm orm.Ormer
	name string
}

var (
	AdminDaoEntity *AdminDao
)

func (dao *AdminDao) Init(name string) {
	AdminDaoEntity = &AdminDao{}
	dao.name = name
	o := common.GetOrm(name)
	if o == nil {
		panic("orm应该不为空")
	}

	dao.orm = o
}