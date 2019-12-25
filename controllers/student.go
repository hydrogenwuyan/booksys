package controllers

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"project/booksys/common"
	. "project/booksys/error_code"
	"project/booksys/logic"
	"project/booksys/models/dao"
	"strconv"
)

const (
	MaxBorrowNum = 6  // 最大借阅数量
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

	// 反序列化出借阅信息
	info := make([]*dao.BorrowInfo, 0, 8)
	err = json.Unmarshal([]byte(stuEntity.BorrowInfo), &info)
	if err != nil {
		common.LogFuncError("unmarshal error: %v", err)
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	type respMsg struct {
		Id         string            `json:"id"`
		User       string            `json:"user"`
		Sex        int8              `json:"sex"`
		Age        int8              `json:"age"`
		Phone      string            `json:"phone"`
		Name       string            `json:"name"`
		BorrowInfo []*dao.BorrowInfo `json:"borrowInfo"`
	}

	resp := &respMsg{
		Id:         fmt.Sprintf("%d", stuEntity.Id),
		User:       stuEntity.User,
		Sex:        stuEntity.Sex,
		Age:        stuEntity.Age,
		Phone:      stuEntity.Phone,
		Name:       stuEntity.Name,
		BorrowInfo: info,
	}

	c.SuccessResponse(resp)
}

// 填写信息
func (c *StudentControllers) MyInfo() {

}

// 借书
func (c *StudentControllers) Borrow() {
	stuId, errCode := c.ParseToken()
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

}
