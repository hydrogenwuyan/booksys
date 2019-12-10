package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	. "project/booksys/error_code"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) GetPost(reqData interface{}) (err error) {
	err = json.Unmarshal(c.Ctx.Input.RequestBody, reqData)
	if err != nil {
		logs.Error("BaseController GetPost Unmarshal Fail, error: ", err)
		return
	}
	return
}

// 不带数据成功返回
func (c *BaseController) SuccessResponseWithoutData() {
	err := c.Ctx.Output.Body([]byte(fmt.Sprintf("{\"code\":%d}", ERROR_CODE_SUCCESS)))
	if err != nil {
		logs.Error("BaseController SuccessResponseWithoutData, error: ", err)
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
		logs.Error("BaseController SuccessResponse, error: ", err)
	}
}

//错误返回
func (c *BaseController) ErrorResponse(errCode ERROR_CODE) {
	msg := errCode.String()
	err := c.Ctx.Output.Body([]byte(fmt.Sprintf("{\"code\":%d, \"msg\":\"%s\"}", errCode, msg)))
	if err != nil {
		logs.Error("BaseController SuccessResponse, error: ", err)
	}
}
