package dao

import (
	"fmt"
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
	e.CreateTime = timeutils.Now()
	_, err = dao.orm.Insert(e)
	if err != nil {
		common.LogFuncError("adminDao create, error: %v", err)
		return
	}

	return
}

func (dao *BookDao) Info(page, limit int32, name string) (entityList []*entity.BookEntity, err error) {
	offset := (page - 1) * limit
	entityList = make([]*entity.BookEntity, 0, 8)
	sql := fmt.Sprintf("select %s,%s,%s,%s,%s,%s,%s,%s,%s from %s where %s like ? order by %s limit ?,?",
		entity.COLUMN_BookEntity_Id,
		entity.COLUMN_BookEntity_IsBorrow,
		entity.COLUMN_BookEntity_Type,
		entity.COLUMN_BookEntity_Name,
		entity.COLUMN_BookEntity_Author,
		entity.COLUMN_BookEntity_Press,
		entity.COLUMN_BookEntity_CreateTime,
		entity.COLUMN_BookEntity_UpdateTime,
		entity.COLUMN_BookEntity_DeleteTime,
		entity.TABLE_BookEntity,
		entity.COLUMN_BookEntity_Name,
		entity.COLUMN_BookEntity_Id)
	name = "%" + name + "%"
	_, err = dao.orm.Raw(sql, name, offset, limit).QueryRows(&entityList)
	if err != nil {
		common.LogFuncError("BookDao info, error: %v", err)
		return
	}

	return
}

func (dao *BookDao) FetchByAuthor(page, limit int32, author string) (entityList []*entity.BookEntity, err error) {
	offset := (page - 1) * limit
	entityList = make([]*entity.BookEntity, 0, 8)
	sql := fmt.Sprintf("select %s,%s,%s,%s,%s,%s,%s,%s,%s from %s where %s like ? order by %s limit ?,?",
		entity.COLUMN_BookEntity_Id,
		entity.COLUMN_BookEntity_IsBorrow,
		entity.COLUMN_BookEntity_Type,
		entity.COLUMN_BookEntity_Name,
		entity.COLUMN_BookEntity_Author,
		entity.COLUMN_BookEntity_Press,
		entity.COLUMN_BookEntity_CreateTime,
		entity.COLUMN_BookEntity_UpdateTime,
		entity.COLUMN_BookEntity_DeleteTime,
		entity.TABLE_BookEntity,
		entity.COLUMN_BookEntity_Author,
		entity.COLUMN_BookEntity_Id)
	author = "%" + author + "%"
	_, err = dao.orm.Raw(sql, author, offset, limit).QueryRows(&entityList)
	if err != nil {
		common.LogFuncError("BookDao fetchByAuthor, error: %v", err)
		return
	}

	return
}

func (dao *BookDao) FetchByType(page, limit int32, typ int32) (entityList []*entity.BookEntity, err error) {
	offset := (page - 1) * limit
	entityList = make([]*entity.BookEntity, 0, 8)
	sql := fmt.Sprintf("select %s,%s,%s,%s,%s,%s,%s,%s,%s from %s where %s=? order by %s limit ?,?",
		entity.COLUMN_BookEntity_Id,
		entity.COLUMN_BookEntity_IsBorrow,
		entity.COLUMN_BookEntity_Type,
		entity.COLUMN_BookEntity_Name,
		entity.COLUMN_BookEntity_Author,
		entity.COLUMN_BookEntity_Press,
		entity.COLUMN_BookEntity_CreateTime,
		entity.COLUMN_BookEntity_UpdateTime,
		entity.COLUMN_BookEntity_DeleteTime,
		entity.TABLE_BookEntity,
		entity.COLUMN_BookEntity_Type,
		entity.COLUMN_BookEntity_Id)
	_, err = dao.orm.Raw(sql, typ, offset, limit).QueryRows(&entityList)
	if err != nil {
		common.LogFuncError("BookDao fetchByAuthor, error: %v", err)
		return
	}

	return
}
