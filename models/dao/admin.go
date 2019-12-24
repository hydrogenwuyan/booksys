package dao

import (
	"github.com/astaxie/beego/orm"
	"project/booksys/common"
	"project/booksys/models/entity"
)

type SexType int8

const (
	SexTypeMan SexType = iota
	SexTypeWoman
)

func (t SexType) Vaild() bool {
	switch t {
	case SexTypeMan,
		SexTypeWoman:
		return true
	default:
		return false
	}
}

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

func (dao *AdminDao) Info(id int64) (e *entity.AdminEntity, err error) {
	e = &entity.AdminEntity{}
	err = AdminDaoEntity.orm.QueryTable(entity.TABLE_AdminEntity).Filter(entity.COLUMN_AdminEntity_Id, id).One(e)
	if err != nil {
		if err == orm.ErrNoRows {
			err = nil
			e.Id = 0
			return
		}
		common.LogFuncError("AdminDao Fetch, error: %v", err)
		return
	}

	return
}

func (dao *AdminDao) Fetch(user string, pass string) (e *entity.AdminEntity, err error) {
	e = &entity.AdminEntity{}
	err = AdminDaoEntity.orm.QueryTable(entity.TABLE_AdminEntity).Filter(entity.COLUMN_AdminEntity_User, user).Filter(entity.COLUMN_AdminEntity_Password, pass).One(e)
	if err != nil {
		if err == orm.ErrNoRows {
			err = nil
			e.Id = 0
			return
		}
		common.LogFuncError("AdminDao Fetch, error: %v", err)
		return
	}

	return
}

func (dao *AdminDao) Update(e *entity.AdminEntity) (err error) {
	_, err = AdminDaoEntity.orm.Update(e)
	if err != nil {
		common.LogFuncError("AdminDao Fetch, error: %v", err)
		return
	}

	return
}
