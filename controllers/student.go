package controllers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"project/booksys/common"
	. "project/booksys/error_code"
	"project/booksys/logic"
	"project/booksys/models/dao"
	"strconv"
)

const (
	MaxBorrowDay = 30 // 最大借阅天数
)

// 学生
type StudentControllers struct {
	BaseController
}

// 注册
func (c *StudentControllers) Register() {
	type ReqMsg struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}

	reqMsg := &ReqMsg{}
	err := c.GetPost(reqMsg)
	if err != nil {
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	// 过滤数据
	if !logic.IsStringOrNum(reqMsg.User) {
		common.LogFuncError("check user error: %v", reqMsg.User)
		c.ErrorResponse(ERROR_CODE_USER_NAME_ERROR)
		return
	}
	if len(reqMsg.User) > UserMaxLen {
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}
	if !logic.IsStringOrNum(reqMsg.Password) {
		common.LogFuncError("check pass error: %v", reqMsg.Password)
		c.ErrorResponse(ERROR_CODE_USER_PASSWORD_ERROR)
		return
	}
	if len(reqMsg.Password) > PassMaxLen {
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	// 判断用户名已被注册
	stuEntity, err := dao.StudentDaoEntity.FetchByUser(reqMsg.User)
	if err != nil {
		c.ErrorResponse(ERROR_CODE_DB_ERROR)
		return
	}
	if stuEntity.Id != 0 {
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	// 加密
	pass, err := bcrypt.GenerateFromPassword([]byte(reqMsg.Password), bcrypt.DefaultCost)
	if err != nil {
		common.LogFuncError("check pass error: %v", reqMsg.Password)
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	// 创建学生数据
	err = dao.StudentDaoEntity.Create(reqMsg.User, string(pass))
	if err != nil {
		c.ErrorResponse(ERROR_CODE_DB_ERROR)
		return
	}

	c.SuccessResponseWithoutData()
}

// 登陆
func (c *StudentControllers) Login() {
	type ReqMsg struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}

	reqMsg := &ReqMsg{}
	err := c.GetPost(reqMsg)
	if err != nil {
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	// 过滤数据
	if !logic.IsStringOrNum(reqMsg.User) {
		common.LogFuncError("check user error: %v", reqMsg.User)
		c.ErrorResponse(ERROR_CODE_USER_NAME_ERROR)
		return
	}
	if len(reqMsg.User) > UserMaxLen {
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}
	if !logic.IsStringOrNum(reqMsg.Password) {
		common.LogFuncError("check pass error: %v", reqMsg.Password)
		c.ErrorResponse(ERROR_CODE_USER_PASSWORD_ERROR)
		return
	}
	if len(reqMsg.Password) > PassMaxLen {
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	// 查询数据库验证密码
	stuEntity, err := dao.StudentDaoEntity.FetchByUser(reqMsg.User)
	if err != nil {
		c.ErrorResponse(ERROR_CODE_DB_ERROR)
		return
	}
	if stuEntity.Id == 0 {
		c.ErrorResponse(ERROR_CODE_USER_PASSWORD_ERROR)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(stuEntity.Password), []byte(reqMsg.Password))
	if err != nil {
		c.ErrorResponse(ERROR_CODE_USER_PASSWORD_ERROR)
		return
	}

	// 设置token
	errCode := c.SetToken(stuEntity.Id, IdentityStudent)
	if errCode != ERROR_CODE_SUCCESS {
		c.ErrorResponse(errCode)
		return
	}

	type respMsg struct {
		Id    string `json:"id"`
		User  string `json:"user"`
		Sex   int8   `json:"sex"`
		Age   int8   `json:"age"`
		Phone string `json:"phone"`
		Name  string `json:"name"`
	}

	resp := &respMsg{
		Id:    fmt.Sprintf("%d", stuEntity.Id),
		User:  stuEntity.User,
		Sex:   stuEntity.Sex,
		Age:   stuEntity.Age,
		Phone: stuEntity.Phone,
		Name:  stuEntity.Name,
	}

	c.SuccessResponse(resp)
}

// 填写信息
func (c *StudentControllers) MyInfo() {
	id, errCode := c.ParseToken(IdentityStudent)
	if errCode != ERROR_CODE_SUCCESS {
		c.ErrorResponse(errCode)
		return
	}

	type ReqMsg struct {
		Sex   int8   `json:"sex"`
		Age   int8   `json:"age"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}

	reqMsg := &ReqMsg{}
	err := c.GetPost(reqMsg)
	if err != nil {
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	// 验证sex
	if !dao.SexType(reqMsg.Sex).Valid() {
		common.LogFuncWarning("sex warn: %v", reqMsg.Sex)
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	// 验证name
	if !logic.IsStringOrNum(reqMsg.Name) {
		common.LogFuncError("check name error: %v", reqMsg.Name)
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}
	if len(reqMsg.Name) > NameMaxLen {
		common.LogFuncWarning("check name warn: %v", reqMsg.Name)
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	// 验证phone
	if !logic.IsPhone(reqMsg.Phone) {
		common.LogFuncError("check phone error: %v", reqMsg.Phone)
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}
	if len(reqMsg.Phone) > PhoneMaxLen {
		common.LogFuncWarning("check phone warn: %v", reqMsg.Phone)
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	entity, err := dao.StudentDaoEntity.Info(id)
	if err != nil || entity.Id == 0 {
		c.ErrorResponse(ERROR_CODE_DB_ERROR)
		return
	}

	// 更新数据
	entity.Sex = reqMsg.Sex
	entity.Age = reqMsg.Age
	entity.Name = reqMsg.Name
	entity.Phone = reqMsg.Phone
	err = dao.StudentDaoEntity.Update(entity)
	if err != nil {
		c.ErrorResponse(ERROR_CODE_DB_ERROR)
		return
	}

	c.SuccessResponseWithoutData()
}

// 借书
func (c *StudentControllers) Borrow() {
	stuId, errCode := c.ParseToken(IdentityStudent)
	if errCode != ERROR_CODE_SUCCESS {
		c.ErrorResponse(errCode)
		return
	}

	type ReqMsg struct {
		BookId string `json:"bookId"`
		Day    int8   `json:"day"`
	}

	reqMsg := &ReqMsg{}
	err := c.GetPost(reqMsg)
	if err != nil {
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	bookId, err := strconv.ParseInt(reqMsg.BookId, 10, 64)
	if err != nil {
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	// 超出最大借阅天数
	if reqMsg.Day > MaxBorrowDay {
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	// 借书
	errCode = logic.BorrowBook(stuId, bookId, reqMsg.Day)
	if errCode != ERROR_CODE_SUCCESS {
		c.ErrorResponse(errCode)
		return
	}

	c.SuccessResponseWithoutData()
}

// 还书
func (c *StudentControllers) GiveBack() {
	stuId, errCode := c.ParseToken(IdentityStudent)
	if errCode != ERROR_CODE_SUCCESS {
		c.ErrorResponse(errCode)
		return
	}

	type ReqMsg struct {
		BookId string `json:"bookId"`
	}

	reqMsg := &ReqMsg{}
	err := c.GetPost(reqMsg)
	if err != nil {
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	bookId, err := strconv.ParseInt(reqMsg.BookId, 10, 64)
	if err != nil {
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	// 还书
	errCode = logic.GiveBackBook(stuId, bookId)
	if errCode != ERROR_CODE_SUCCESS {
		c.ErrorResponse(errCode)
		return
	}

	c.SuccessResponseWithoutData()
}

func (c *StudentControllers) BorrowInfo() {
	stuId, errCode := c.ParseToken(IdentityStudent)
	if errCode != ERROR_CODE_SUCCESS {
		c.ErrorResponse(errCode)
		return
	}

	_, bookList, err := dao.BookDaoEntity.FetchByStuId(stuId)
	if err != nil {
		return
	}

	c.SuccessResponse(bookList)
}
