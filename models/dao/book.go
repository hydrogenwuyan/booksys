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

const (
	BookNotIsBorrow = iota // 未被借阅
	BookIsBorrow           // 已被借阅
)

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
		common.LogFuncError("BookDao create, error: %v", err)
		return
	}

	return
}

func (dao *BookDao) Info(page, limit int32, name string) (entityList []*entity.BookEntity, err error) {
	offset := (page - 1) * limit
	entityList = make([]*entity.BookEntity, 0, 8)
	sql := fmt.Sprintf("select %s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s from %s where %s like ? order by %s limit ?,?",
		entity.COLUMN_BookEntity_Id,
		entity.COLUMN_BookEntity_IsBorrow,
		entity.COLUMN_BookEntity_StuId,
		entity.COLUMN_BookEntity_StuUser,
		entity.COLUMN_BookEntity_ExpireTime,
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
	sql := fmt.Sprintf("select %s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s from %s where %s like ? order by %s limit ?,?",
		entity.COLUMN_BookEntity_Id,
		entity.COLUMN_BookEntity_IsBorrow,
		entity.COLUMN_BookEntity_StuId,
		entity.COLUMN_BookEntity_StuUser,
		entity.COLUMN_BookEntity_ExpireTime,
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
	sql := fmt.Sprintf("select %s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s from %s where %s=? order by %s limit ?,?",
		entity.COLUMN_BookEntity_Id,
		entity.COLUMN_BookEntity_IsBorrow,
		entity.COLUMN_BookEntity_StuId,
		entity.COLUMN_BookEntity_StuUser,
		entity.COLUMN_BookEntity_ExpireTime,
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

func (dao *BookDao) FetchByIdAndIsBorrow(id int64, isBorrow int8) (e *entity.BookEntity, err error) {
	e = &entity.BookEntity{}
	err = dao.orm.QueryTable(entity.TABLE_BookEntity).Filter(entity.COLUMN_BookEntity_Id, id).Filter(entity.COLUMN_BookEntity_IsBorrow, isBorrow).One(e)
	if err != nil {
		if err == orm.ErrNoRows {
			err = nil
			e.Id = 0
		}
		common.LogFuncError("BookDao fetchByIdAndIsBorrow, error: %v", err)
		return
	}

	return
}

func (dao *BookDao) FetchExpireStuIdList() (stuIdList []int64, err error) {
	stuIdList = make([]int64, 0, 8)
	sql := fmt.Sprintf("select %s from %s where %s<? and %s<>0", entity.COLUMN_BookEntity_StuId, entity.TABLE_BookEntity, entity.COLUMN_BookEntity_ExpireTime, entity.COLUMN_BookEntity_StuId)
	err = dao.orm.Raw(sql, timeutils.Now()).QueryRow(&stuIdList)
	if err != nil {
		if err == orm.ErrNoRows {
			err = nil
			return
		}
		common.LogFuncError("BookDao fetchExpireStuIdList, error: %v", err)
		return
	}

	return
}

func (dao *BookDao) FetchByStuId(stuId int64) (num int64, bookList []*entity.BookEntity, err error) {
	bookList = make([]*entity.BookEntity, 0, 8)
	num, err = dao.orm.QueryTable(entity.TABLE_BookEntity).Filter(entity.COLUMN_BookEntity_StuId, stuId).All(&bookList)
	if err != nil {
		if err == orm.ErrNoRows {
			err = nil
			return
		}
		common.LogFuncError("BookDao fetchExpireStuIdList, error: %v", err)
		return
	}

	return
}

func (dao *BookDao) FetchByStuUser(stuUser string) (num int64, bookList []*entity.BookEntity, err error) {
	bookList = make([]*entity.BookEntity, 0, 8)
	num, err = dao.orm.QueryTable(entity.TABLE_BookEntity).Filter(entity.COLUMN_BookEntity_StuUser, stuUser).All(&bookList)
	if err != nil {
		if err == orm.ErrNoRows {
			err = nil
			return
		}
		common.LogFuncError("BookDao fetchExpireStuIdList, error: %v", err)
		return
	}

	return
}

func (dao *BookDao) Update(e *entity.BookEntity) (err error) {
	_, err = dao.orm.Update(e)
	if err != nil {
		common.LogFuncError("BookDao update, error: %v", err)
		return
	}

	return
}
