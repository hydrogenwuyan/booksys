package controllers

import (
	. "project/booksys/error_code"
	"project/booksys/models/dao"
)

// 图书
type BookControllers struct {
	BaseController
}

// 根据图书名查询
func (c *BookControllers) Name() {
	page, limit, err := c.GetReqPage()
	if err != nil {
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	name := c.GetString("name")
	// TODO: 正则
	list, err := dao.BookDaoEntity.Info(page, limit, name)
	if err != nil {
		c.ErrorResponse(ERROR_CODE_DB_ERROR)
		return
	}

	c.SuccessResponse(list)
}

// 根据作者查询
func (c *BookControllers) Author() {
	page, limit, err := c.GetReqPage()
	if err != nil {
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	author := c.GetString("author")
	// TODO: 正则
	list, err := dao.BookDaoEntity.FetchByAuthor(page, limit, author)
	if err != nil {
		c.ErrorResponse(ERROR_CODE_DB_ERROR)
		return
	}

	c.SuccessResponse(list)
}

// 根据类型查询
func (c *BookControllers) BookType() {
	page, limit, err := c.GetReqPage()
	if err != nil {
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	typ, err := c.GetInt32("type")
	if err != nil {
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	// 检查type
	if !dao.BookType(typ).Valid() {
		c.ErrorResponse(ERROR_CODE_ERROR)
		return
	}

	list, err := dao.BookDaoEntity.FetchByType(page, limit, typ)
	if err != nil {
		c.ErrorResponse(ERROR_CODE_DB_ERROR)
		return
	}

	c.SuccessResponse(list)
}
