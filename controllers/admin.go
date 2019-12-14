package controllers

import (
	"fmt"
	"project/booksys/common"
	. "project/booksys/error_code"
	"project/booksys/models/dao"
	"regexp"
)

// 管理员
type AdminControllers struct {
	BaseController
}

type adminControllersLoginReq struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// 登陆
func (c *AdminControllers) Login() {
	msg := &adminControllersLoginReq{}
	err := c.GetPost(msg)
	if err != nil {
		fmt.Println(*msg)
		return
	}

	// 过滤数据
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	userErr := reg.FindAllString(msg.User, -1)
	if len(userErr) != 0 {
		common.LogFuncError("userErr: %v", userErr)
		c.ErrorResponse(ERROR_CODE_USER_NAME_ERROR)
		return
	}
	passErr := reg.FindAllString(msg.Password, -1)
	if len(passErr) != 0 {
		common.LogFuncError("passErr: %v", passErr)
		c.ErrorResponse(ERROR_CODE_USER_PASSWORD_ERROR)
		return
	}

	// 查询数据库验证密码
	adminEntity, err := dao.AdminDaoEntity.Fetch(msg.User, msg.Password)
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
		Id    int64  `json:"id"`
		User  string `json:"user"`
		Sex   int8   `json:"sex"`
		Age   int8  `json:"age"`
		Phone string  `json:"phone"`
		Name  string `json:"name"`
	}

	resp := &respMsg{
		Id:    adminEntity.Id,
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

}