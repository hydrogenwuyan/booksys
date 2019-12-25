package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"project/booksys/common"
	. "project/booksys/error_code"
	"project/booksys/utils/tokenutils"
)

const (
	TokenKey = "token"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) GetReqPage() (page, limit int32, err error) {
	pageInt, err := c.GetInt("page", 1)
	if err != nil {
		common.LogFuncWarning("get parameter %s failed : %v", page, err)
		return
	}
	if pageInt < 1 {
		pageInt = 1
	}

	limitInt, err := c.GetInt("limit", 10)
	if err != nil {
		common.LogFuncWarning("get parameter %s failed : %v", limit, err)
		return
	}
	if limitInt > 10 {
		limitInt = 10
	}

	page = int32(pageInt)
	limit = int32(limitInt)

	return
}

func (c *BaseController) GetPost(reqData interface{}) (err error) {
	err = json.Unmarshal(c.Ctx.Input.RequestBody, reqData)
	if err != nil {
		common.LogFuncError("BaseController GetPost Unmarshal Fail, error: %v", err)
		return
	}
	return
}

// 不带数据成功返回
func (c *BaseController) SuccessResponseWithoutData() {
	err := c.Ctx.Output.Body([]byte(fmt.Sprintf("{\"code\":%d}", ERROR_CODE_SUCCESS)))
	if err != nil {
		common.LogFuncError("BaseController SuccessResponseWithoutData Fail, error: %v", err)
	}
}

// 带数据成功返回
func (c *BaseController) SuccessResponse(result interface{}) {

	params := map[string]interface{}{
		"code": ERROR_CODE_SUCCESS,
		"data": result,
	}

	err := c.Ctx.Output.JSON(params, false, false)
	if err != nil {
		common.LogFuncError("BaseController SuccessResponse Fail, error: %v", err)
	}
}

//错误返回
func (c *BaseController) ErrorResponse(errCode ERROR_CODE) {
	msg := errCode.String()
	err := c.Ctx.Output.Body([]byte(fmt.Sprintf("{\"code\":%d, \"msg\":\"%s\"}", errCode, msg)))
	if err != nil {
		common.LogFuncError("BaseController ErrorResponse Fail, error: %v", err)
	}
}

// 设置token
func (c *BaseController) SetToken(id int64) (errCode ERROR_CODE) {
	errCode = ERROR_CODE_SUCCESS
	token, err := tokenutils.GenerateToken(id)
	if err != nil {
		errCode = ERROR_CODE_GENERATE_TOKEN_FAIL
		return
	}

	err = tokenutils.SetToken(id, token)
	if err != nil {
		errCode = ERROR_CODE_SET_TOKEN_FAIL
		return
	}

	c.Ctx.SetCookie(TokenKey, token, tokenutils.AccessTokenExpiredSecs)
	return
}

// 解析token
func (c *BaseController) ParseToken() (id int64, errCode ERROR_CODE) {
	errCode = ERROR_CODE_SUCCESS
	token := c.Ctx.GetCookie(TokenKey)
	result, id := tokenutils.CheckAndParseToken(token)
	if result != tokenutils.TokenOk {
		errCode = ERROR_CODE_TOKEN_EXPIRED
		return
	}

	return
}

// 清理cookie
func (c *BaseController) ClearCookieToken() {
	ok := tokenutils.ClearToken(c.Ctx.GetCookie(TokenKey))
	if !ok {
		common.LogFuncError("ClearToken Fail")
	}

	c.Ctx.SetCookie(TokenKey, "", -1)
}

// 清理token
func (c *BaseController) ClearToken() {
	token := c.Ctx.GetCookie(TokenKey)
	ok := tokenutils.ClearToken(token)
	if !ok {
		common.LogFuncError("ClearToken Fail")
	}
}
