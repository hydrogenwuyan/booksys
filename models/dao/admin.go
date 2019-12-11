package dao

import (
	"github.com/astaxie/beego/orm"
	"github.com/siddontang/go/log"
	"project/booksys/common"
	"project/booksys/models/entity"
)

type AdminDao struct {
	orm  orm.Ormer
	name string
}

var (
	AdminDaoEntity *AdminDao
)

func NewAdminDao(name string) (dao *AdminDao) {
	dao = &AdminDao{}
	dao.name = name
	o := common.GetOrm(name)
	if o == nil {
		panic("orm应该不为空")
	}
	
	dao.orm = o
	return
}

func (dao *AdminDao) Fetch(user string, pass string) (e *entity.AdminEntity, err error) {
	e = &entity.AdminEntity{}
	err = AdminDaoEntity.orm.QueryTable(entity.TABLE_AdminEntity).Filter(entity.COLUMN_AdminEntity_User, user).Filter(entity.COLUMN_AdminEntity_Password, pass).One(e)
	if err != nil {
		if err == orm.ErrNoRows {
			err = nil
			return
		}
		log.Error("AdminDao Fetch, error: ", err)
		return
	}

	return
}
