package controllers

import (
	"fmt"
	. "project/booksys/error_code"
	"project/booksys/models/dao"
	"project/booksys/utils/tokenutils"
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

	// 设置token
	flag := false
	status := tokenutils.CheckTokenSignature(fmt.Sprint(adminEntity.Id))
	if status == tokenutils.TokenOk {
		flag = true
	}

	var token string
	if !flag {
		token, err = tokenutils.GenerateToken(adminEntity.Id)
		if err != nil {
			c.ErrorResponse(ERROR_CODE_GENERATE_TOKEN_FAIL)
			return
		}
	}

	c.Ctx.SetCookie(TokenKey, token, tokenutils.AccessTokenExpiredSecs)

	type respMsg struct {
		Id    int64  `json:"id"`
		User  string `json:"user"`
		Sex   int8   `json:"sex"`
		Age   int32  `json:"age"`
		Phone int32  `json:"phone"`
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
