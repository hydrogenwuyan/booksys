package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"project/booksys/common"
	"project/booksys/models/entity"
	"project/booksys/utils/timeutils"
)

const (
	StudentNotIsBlack = iota // 不在黑名单
	StudentIsBlack           // 黑名单
)

type StudentDao struct {
	orm  orm.Ormer
	name string
}

var (
	StudentDaoEntity *StudentDao
)

func NewStudentDao(name string) (dao *StudentDao) {
	dao = &StudentDao{}
	dao.name = name
	o := common.GetOrm(name)
	if o == nil {
		panic("orm应该不为空")
	}

	dao.orm = o
	return
}

func (dao *StudentDao) Create(user string, pass string) (err error) {
	stu := &entity.StudentEntity{
		User:       user,
		Password:   pass,
		CreateTime: timeutils.Now(),
	}
	_, err = dao.orm.Insert(stu)
	if err != nil {
		common.LogFuncError("studentDao create, error: %v", err)
		return
	}

	return
}

func (dao *StudentDao) FetchByUser(user string) (e *entity.StudentEntity, err error) {
	e = &entity.StudentEntity{}
	err = dao.orm.QueryTable(entity.TABLE_StudentEntity).Filter(entity.COLUMN_StudentEntity_User, user).One(e)
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

func (dao *StudentDao) Info(id int64) (e *entity.StudentEntity, err error) {
	e = &entity.StudentEntity{}
	err = dao.orm.QueryTable(entity.TABLE_StudentEntity).Filter(entity.COLUMN_StudentEntity_Id, id).One(e)
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

func (dao *StudentDao) List() (list []*entity.StudentEntity, err error) {
	list = make([]*entity.StudentEntity, 0, 8)
	_, err = dao.orm.QueryTable(entity.TABLE_StudentEntity).Filter(entity.COLUMN_StudentEntity_IsBlack, StudentIsBlack).All(&list)
	if err != nil {
		if err == orm.ErrNoRows {
			err = nil
			return
		}
		common.LogFuncError("adminDao fetch, error: %v", err)
		return
	}

	return
}

func (dao *StudentDao) UpdateByUser(user string) (err error) {
	sql := fmt.Sprintf("update %s set %s=? where %s=?", entity.TABLE_StudentEntity, entity.COLUMN_StudentEntity_IsBlack, entity.COLUMN_StudentEntity_User)
	_, err = dao.orm.Raw(sql, StudentNotIsBlack, user).Exec()
	if err != nil {
		if err == orm.ErrNoRows {
			err = nil
			return
		}
		common.LogFuncError("adminDao fetch, error: %v", err)
		return
	}

	return
}

func (dao *StudentDao) Update(e *entity.StudentEntity) (err error) {
	e.UpdateTime = timeutils.Now()
	_, err = dao.orm.Update(e)
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
