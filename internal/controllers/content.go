package controllers

import (
	"net/http"
	"strconv"

	"github.com/MNU/exam-go"
	"github.com/labstack/echo"
)

// ContentController is
type ContentController struct {
	contentSvc goexam.ContentService
}

// NewContentController is
func NewContentController(contentSvc goexam.ContentService) *ContentController {
	return &ContentController{
		contentSvc,
	}
}

// Create 创建新的考试
func (c *ContentController) Create(ctx echo.Context) error {
	contentRequest := new(goexam.Content)
	err := ctx.Bind(contentRequest)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	content, err := c.contentSvc.Create(contentRequest)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, content)
}

// Delete is 删除考试
func (c *ContentController) Delete(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	if err != nil || id == 0 {
		return ctx.String(http.StatusBadRequest, "参数不全")
	}
	err = c.contentSvc.Delete(uint(id))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "服务器错误")
	}
	return ctx.NoContent(http.StatusOK)
}

// Update is
func (c *ContentController) Update(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	if err != nil || id == 0 {
		return ctx.String(http.StatusBadRequest, "参数不全")
	}
	contentRequest := new(goexam.Content)
	err = ctx.Bind(contentRequest)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "参数不全")
	}
	contentRequest.ID = uint(id)
	err = c.contentSvc.Update(contentRequest)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "server")
	}
	return ctx.NoContent(http.StatusOK)
}

// Get is
func (c *ContentController) Get(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	if err != nil || id == 0 {
		return ctx.String(http.StatusBadRequest, "参数不全")
	}
	content, err := c.contentSvc.Get(uint(id))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "服务错误")
	}
	return ctx.JSON(http.StatusOK, content)
}

// GetList is
func (c *ContentController) GetList(ctx echo.Context) error {
	filter := new(goexam.ContentFilter)
	err := ctx.Bind(filter)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "参数不全")
	}
	contents, err := c.contentSvc.GetList(filter)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, contents)
}
