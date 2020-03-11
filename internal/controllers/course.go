package controllers

import (
	"net/http"
	"strconv"

	"github.com/goexam"
	"github.com/labstack/echo"
)

// CourseController is
type CourseController struct {
	courseSvc goexam.CourseService
}

// NewCourseController is
func NewCourseController(courseSvc goexam.CourseService) *CourseController {
	return &CourseController{
		courseSvc,
	}
}

// Create is
func (c *CourseController) Create(ctx echo.Context) error {
	course := new(goexam.Course)
	err := ctx.Bind(course)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "参数错误")
	}
	err = c.courseSvc.Create(course)
	return ctx.NoContent(http.StatusOK)
}

// Update is
func (c *CourseController) Update(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	if id == 0 || err != nil {
		return ctx.JSON(http.StatusBadRequest, "参数错误")
	}
	course := new(goexam.Course)
	err = ctx.Bind(course)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "参数错误")
	}
	course.ID = uint(id)
	err = c.courseSvc.Update(course)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "服务器错误")
	}
	return ctx.NoContent(http.StatusOK)
}

// Delele is
func (c *CourseController) Delele(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	err = c.courseSvc.Delete(uint(id))
	if err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	return ctx.NoContent(http.StatusOK)
}

// Get is ...
func (c *CourseController) Get(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	course, err := c.courseSvc.Get(uint(id))
	if err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	return ctx.JSON(http.StatusOK, course)
}

// GetList is
func (c *CourseController) GetList(ctx echo.Context) error {
	filter := new(goexam.CourseFilter)
	err := ctx.Bind(filter)
	filter.LoadDefault()
	if err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	courseList, err := c.courseSvc.GetList(filter)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "服务错误")
	}
	return ctx.JSON(http.StatusOK, courseList)
}
