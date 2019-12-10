package controllers

import (
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
	c.GetPost(msg)

	// 过滤数据
	reg := regexp.MustCompile(`^[a-zA-Z0-9]`)
	userErr := reg.FindAllString(msg.User, -1)
	if len(userErr) != 0 {
		c.ErrorResponse(ERROR_CODE_USER_NAME_ERROR)
		return
	}
	passErr := reg.FindAllString(msg.Password, -1)
	if len(passErr) != 0 {
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

	c.SuccessResponse(adminEntity)
}
