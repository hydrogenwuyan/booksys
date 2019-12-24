package dao

import (
	"github.com/astaxie/beego/orm"
	"project/booksys/common"
	"project/booksys/models/entity"
	"project/booksys/utils/timeutils"
)

type BookType int32

const (
	BookTypeChinese BookType = iota // 简写
	BookTypeForeignLanguage
)

func (t BookType) Valid() bool {
	switch t {
	case BookTypeChinese,
		BookTypeForeignLanguage:
		return true
	default:
		return false
	}
}

type BookDao struct {
	orm  orm.Ormer
	name string
}

var (
	BookDaoEntity *BookDao
)

func NewBookDao(name string) (dao *BookDao) {
	dao = &BookDao{}
	dao.name = name
	o := common.GetOrm(name)
	if o == nil {
		panic("orm应该不为空")
	}

	dao.orm = o
	return
}

func (dao *BookDao) Create(e *entity.BookEntity) (err error) {
	e.UpdateTime = timeutils.Now()
	_, err = dao.orm.Insert(e)
	if err != nil {
		common.LogFuncError("adminDao create, error: %v", err)
		return
	}

	return
}
