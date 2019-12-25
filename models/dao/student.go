package dao

import (
	"github.com/astaxie/beego/orm"
	"project/booksys/common"
	"project/booksys/models/entity"
	"project/booksys/utils/timeutils"
)

// 借阅信息
type BorrowInfo struct {
	Id        int64  `json:"book"`
	StartTime int64  `json:"startTime"` // 借阅时间
	Day       int8   `json:"day"`       // 最晚归还时间
	Name      string `json:"name"`      // 书名
	Author    string `json:"author"`    // 作者
}

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

func (dao *StudentDao) Create(user string, pass string, data string) (err error) {
	stu := &entity.StudentEntity{
		User:       user,
		Password:   pass,
		BorrowInfo: data,
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
