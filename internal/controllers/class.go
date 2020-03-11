package controllers

import (
	"net/http"
	"strconv"

	"github.com/goexam"
	"github.com/labstack/echo"
)

// ClassController is
type ClassController struct {
	classSvc goexam.ClassService
}

// NewClassController is
func NewClassController(classSvc goexam.ClassService) *ClassController {
	return &ClassController{
		classSvc,
	}
}

// Create is 添加 class
func (c *ClassController) Create(ctx echo.Context) error {
	class := new(goexam.Class)
	err := ctx.Bind(class)
	err = c.classSvc.Create(class)
	if err != nil {
		return ctx.String(http.StatusBadGateway, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}

// Delete is 删除
func (c *ClassController) Delete(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	if id == 0 || err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}

	err = c.classSvc.Delete(uint(id))
	if err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	return ctx.NoContent(http.StatusOK)
}

// Update is
func (c *ClassController) Update(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	if err != nil || id == 0 {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	class := new(goexam.Class)
	err = ctx.Bind(class)
	class.ID = uint(id)
	err = c.classSvc.Update(class)
	if err != nil {
		return ctx.String(http.StatusBadGateway, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}

// Get is
func (c *ClassController) Get(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	if err != nil || id == 0 {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	class, err := c.classSvc.Get(uint(id))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, class)
}

// GetList is
func (c *ClassController) GetList(ctx echo.Context) error {
	classFilter := new(goexam.ClassFilter)
	err := ctx.Bind(classFilter)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	classList, err := c.classSvc.GetList(classFilter)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	return ctx.JSON(http.StatusOK, classList)
}
