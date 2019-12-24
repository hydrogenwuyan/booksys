package controllers

import (
	"fmt"
	"project/booksys/common"
	. "project/booksys/error_code"
	"project/booksys/logic"
	"project/booksys/models/dao"
)

const (
	NameMaxLen  = 18
	PhoneMaxLen = 18
)

// 管理员
type AdminControllers struct {
	BaseController
}

// 登陆
func (c *AdminControllers) Login() {
	type ReqMsg struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}

	reqMsg := &ReqMsg{}
	err := c.GetPost(reqMsg)
	if err != nil {
		return
	}

	// 过滤数据
	if !logic.IsStringOrNum(reqMsg.User) {
		common.LogFuncError("check user error: %v", reqMsg.User)
		c.ErrorResponse(ERROR_CODE_USER_NAME_ERROR)
		return
	}
	if !logic.IsStringOrNum(reqMsg.Password) {
		common.LogFuncError("check pass error: %v", reqMsg.Password)
		c.ErrorResponse(ERROR_CODE_USER_PASSWORD_ERROR)
		return
	}

	// 查询数据库验证密码
	adminEntity, err := dao.AdminDaoEntity.Fetch(reqMsg.User, reqMsg.Password)
	if err != nil {
		c.ErrorResponse(ERROR_CODE_DB_ERROR)
		return
	}
	if adminEntity.Id == 0 {
		c.ErrorResponse(ERROR_CODE_USER_PASSWORD_ERROR)
		return
	}

	// 设置token
	errCode := c.SetToken(adminEntity.Id)
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
		Id:    fmt.Sprintf("%d", adminEntity.Id),
		User:  adminEntity.User,
		Sex:   adminEntity.Sex,
		Age:   adminEntity.Age,
		Phone: adminEntity.Phone,
		Name:  adminEntity.Name,
	}

	c.SuccessResponse(resp)
}

// 填写个人信息
func (c *AdminControllers) MyInfo() {
	id, errCode := c.ParseToken()
	if errCode != ERROR_CODE_SUCCESS {
		c.ErrorResponse(ERROR_CODE_ERROR)
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
	if !dao.SexType(reqMsg.Sex).Vaild() {
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
	if !logic.IsStringOrNum(reqMsg.Phone) {
		common.LogFuncError("check phone error: %v", reqMsg.Phone)
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}
	if len(reqMsg.Phone) > PhoneMaxLen {
		common.LogFuncWarning("check phone warn: %v", reqMsg.Phone)
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	entity, err := dao.AdminDaoEntity.Info(id)
	if err != nil || entity.Id == 0 {
		c.ErrorResponse(ERROR_CODE_DB_ERROR)
		return
	}

	// 更新数据
	entity.Sex = reqMsg.Sex
	entity.Age = reqMsg.Age
	entity.Name = reqMsg.Name
	entity.Phone = reqMsg.Phone
	err = dao.AdminDaoEntity.Update(entity)
	if err != nil {
		c.ErrorResponse(ERROR_CODE_DB_ERROR)
		return
	}

	c.SuccessResponseWithoutData()
}
