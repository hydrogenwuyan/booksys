package logic

import (
	"fmt"
	"project/booksys/common"
	. "project/booksys/error_code"
	"project/booksys/models/dao"
	"project/booksys/utils/timeutils"
	"regexp"
	"time"
)

const (
	MaxBorrowNum     = 6                // 最大借阅数量
	ExpireBorrowTime = 10 * time.Second // 过期时间
)

func IsStringOrNum(ip string) (b bool) {
	if m, _ := regexp.MatchString("[a-zA-Z0-9]+", ip); !m {
		return false
	}
	return true
}

func IsPhone(phone string) (b bool) {
	if m, _ := regexp.MatchString("[0-9]+", phone); !m {
		return false
	}
	return true
}

// 借书
func BorrowBook(stuId, bookId int64, day int8) (errCode ERROR_CODE) {
	// 查询学生数据
	stuEntity, err := dao.StudentDaoEntity.Info(stuId)
	if err != nil || stuEntity.Id == 0 {
		errCode = ERROR_CODE_DB_ERROR
		return
	}

	// 黑名单判断
	if stuEntity.IsBlack != dao.StudentNotIsBlack {
		errCode = ERROR_CODE_DB_ERROR
		return
	}

	// 判断是否到了借阅上限
	num, _, err := dao.BookDaoEntity.FetchByStuId(stuId)
	if err != nil {
		errCode = ERROR_CODE_DB_ERROR
		return
	}
	if num >= MaxBorrowNum {
		errCode = ERROR_CODE_ERROR
		return
	}

	// 取图书数据
	bookEntity, err := dao.BookDaoEntity.FetchByIdAndIsBorrow(bookId, dao.BookNotIsBorrow)
	if err != nil {
		errCode = ERROR_CODE_DB_ERROR
		return
	}
	if bookEntity.Id == 0 {
		errCode = ERROR_CODE_BOOK_NOT_EXIST
		return
	}

	// 加锁
	bookKey := fmt.Sprintf("borrow.%d", bookId)
	val := common.RedisClient.Incr(bookKey).Val()
	if val != 1 {
		common.LogFuncDebug("redis incr not equal 1, val: %v", val)
		errCode = ERROR_CODE_ERROR
		return
	}

	// 设置过期时间
	err = common.RedisClient.Expire(bookKey, ExpireBorrowTime).Err()
	if err != nil {
		common.LogFuncError("redis expire fail, err: %v", err)
	}
	defer func() {
		common.RedisClient.Del(bookKey)
	}()

	// 更新图书信息
	bookEntity.ExpireTime = timeutils.Now() + int64(time.Hour/time.Millisecond)*24*int64(day)
	bookEntity.StuId = stuId
	bookEntity.StuUser = stuEntity.User
	bookEntity.IsBorrow = dao.BookIsBorrow
	err = dao.BookDaoEntity.Update(bookEntity)
	if err != nil {
		errCode = ERROR_CODE_DB_ERROR
		return
	}

	errCode = ERROR_CODE_SUCCESS
	return
}

// 还书
func GiveBackBook(stuId, bookId int64) (errCode ERROR_CODE) {
	// 查询学生数据
	stuEntity, err := dao.StudentDaoEntity.Info(stuId)
	if err != nil || stuEntity.Id == 0 {
		errCode = ERROR_CODE_DB_ERROR
		return
	}

	// 黑名单无法还书
	if stuEntity.IsBlack != dao.StudentNotIsBlack {
		errCode = ERROR_CODE_DB_ERROR
		return
	}

	bookEntity, err := dao.BookDaoEntity.FetchByIdAndIsBorrow(bookId, dao.BookIsBorrow)
	if err != nil || bookEntity.Id == 0 {
		errCode = ERROR_CODE_DB_ERROR
		return
	}

	if bookEntity.StuId != stuId {
		errCode = ERROR_CODE_ERROR
		return
	}

	// 更新图书数据
	bookEntity.ExpireTime = 0
	bookEntity.StuId = 0
	bookEntity.StuUser = ""
	bookEntity.IsBorrow = dao.BookNotIsBorrow
	err = dao.BookDaoEntity.Update(bookEntity)
	if err != nil {
		errCode = ERROR_CODE_DB_ERROR
		return
	}

	errCode = ERROR_CODE_SUCCESS
	return
}

// 从黑名单删除学生
func DeleteStudentByBlackList(stuUser string) (errCode ERROR_CODE) {

	// 创建事务
	o := common.GetOrm(dao.Database)
	// 事务开始
	err := o.Begin()
	if err != nil {
		common.LogFuncError("DeleteStudentByBlackList begin fail, err: %v", err)
		return
	}
	defer func() {
		if errCode != ERROR_CODE_SUCCESS {
			err = o.Rollback()
			if err != nil {
				common.LogFuncError("DeleteStudentByBlackList roolback fail, err: %v", err)
			}
		}
	}()

	// 更新图书数据
	_, bookList, err := dao.BookDaoEntity.FetchByStuUser(stuUser)
	if err != nil {
		return
	}

	now := timeutils.Now()
	for _, e := range bookList {
		if e.ExpireTime > now {
			continue
		}

		e.ExpireTime = 0
		e.StuId = 0
		e.StuUser = ""
		e.IsBorrow = dao.BookNotIsBorrow
		err = dao.BookDaoEntity.Update(e)
		if err != nil {
			errCode = ERROR_CODE_DB_ERROR
			return
		}
	}

	// 更新学生数据
	err = dao.StudentDaoEntity.UpdateByUser(stuUser)
	if err != nil {
		errCode = ERROR_CODE_DB_ERROR
		return
	}

	// 事务提交
	err = o.Commit()
	if err != nil {
		common.LogFuncError("DeleteStudentByBlackList commit fail, err: %v", err)
	}

	errCode = ERROR_CODE_SUCCESS
	return
}
