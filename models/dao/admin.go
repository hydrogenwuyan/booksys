package dao

import (
	"github.com/astaxie/beego/orm"
	"project/booksys/common"
	"project/booksys/models/entity"
	"project/booksys/utils/timeutils"
)

type SexType int8

const (
	SexTypeMan SexType = iota
	SexTypeWoman
)

func (t SexType) Valid() bool {
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

func (dao *AdminDao) Create(user string, pass string) (err error) {
	admin := &entity.AdminEntity{
		User:       user,
		Password:   pass,
		Phone:      "",
		Name:       "",
		CreateTime: timeutils.Now(),
	}
	_, err = dao.orm.Insert(admin)
	if err != nil {
		common.LogFuncError("adminDao create, error: %v", err)
		return
	}

	return
}

func (dao *AdminDao) Info(id int64) (e *entity.AdminEntity, err error) {
	e = &entity.AdminEntity{}
	err = dao.orm.QueryTable(entity.TABLE_AdminEntity).Filter(entity.COLUMN_AdminEntity_Id, id).One(e)
	if err != nil {
		if err == orm.ErrNoRows {
			err = nil
			e.Id = 0
			return
		}
		common.LogFuncError("adminDao info, error: %v", err)
		return
	}

	return
}

func (dao *AdminDao) FetchByUser(user string) (e *entity.AdminEntity, err error) {
	e = &entity.AdminEntity{}
	err = dao.orm.QueryTable(entity.TABLE_AdminEntity).Filter(entity.COLUMN_AdminEntity_User, user).One(e)
	if err != nil {
		if err == orm.ErrNoRows {
			err = nil
			e.Id = 0
			return
		}
		common.LogFuncError("adminDao fetch, error: %v", err)
		return
	}

	return
}

func (dao *AdminDao) Fetch(user string, pass string) (e *entity.AdminEntity, err error) {
	e = &entity.AdminEntity{}
	err = dao.orm.QueryTable(entity.TABLE_AdminEntity).Filter(entity.COLUMN_AdminEntity_User, user).Filter(entity.COLUMN_AdminEntity_Password, pass).One(e)
	if err != nil {
		if err == orm.ErrNoRows {
			err = nil
			e.Id = 0
			return
		}
		common.LogFuncError("adminDao fetch, error: %v", err)
		return
	}

	return
}

func (dao *AdminDao) Update(e *entity.AdminEntity) (err error) {
	e.UpdateTime = timeutils.Now()
	_, err = dao.orm.Update(e)
	if err != nil {
		common.LogFuncError("adminDao update, error: %v", err)
		return
	}

	return
}
