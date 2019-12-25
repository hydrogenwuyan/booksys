package logic

import (
	"encoding/json"
	"fmt"
	"project/booksys/common"
	. "project/booksys/error_code"
	"project/booksys/models/dao"
	"project/booksys/utils/timeutils"
	"regexp"
	"time"
)

const (
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

func BorrowBook(stuId, bookId int64, day int8) (errCode ERROR_CODE) {
	// 取图书数据
	bookEntity, err := dao.BookDaoEntity.FetchByIdAndIsBorrow(bookId, dao.BookNotIsBorrow)
	if err != nil || bookEntity.Id == 0 {
		errCode = ERROR_CODE_DB_ERROR
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

	// 开启事务
	o := common.GetOrm(dao.Database)
	o.Begin()
	defer func() {
		if err != nil { // 回滚
			o.Rollback()
		}
	}()

	// 更新图书信息
	err = dao.BookDaoEntity.UpdateAboutBorrow(bookId, dao.BookIsBorrow)
	if err != nil {
		errCode = ERROR_CODE_DB_ERROR
		return
	}

	// 查询学生数据
	stuEntity, err := dao.StudentDaoEntity.Info(stuId)
	if err != nil || stuEntity.Id == 0 {
		errCode = ERROR_CODE_DB_ERROR
		return
	}

	// 反序列化出借阅信息
	info := make([]*dao.BorrowInfo, 0, 8)
	err = json.Unmarshal([]byte(stuEntity.BorrowInfo), &info)
	if err != nil {
		common.LogFuncError("unmarshal error: %v", err)
		errCode = ERROR_CODE_ERROR
		return
	}
	info = append(info, &dao.BorrowInfo{
		Id:        bookId,
		StartTime: timeutils.Now(),
		Day:       day,
		Name:      bookEntity.Name,
		Author:    bookEntity.Author,
	})

	data, err := json.Marshal(&info)
	if err != nil {
		common.LogFuncError("marshal error: %v", err)
		errCode = ERROR_CODE_ERROR
		return
	}
	stuEntity.BorrowInfo = string(data)

	// 更新学生信息
	err = dao.StudentDaoEntity.Update(stuEntity)
	if err != nil {
		errCode = ERROR_CODE_DB_ERROR
		return
	}

	// 提交事务
	o.Commit()

	errCode = ERROR_CODE_SUCCESS

	return
}
