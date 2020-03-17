package teacher

import (
	"net/http"
	"strconv"

	"github.com/MNU/exam-go"
	"github.com/labstack/echo"
)

// CollageController is
type CollageController struct {
	collageSvc goexam.CollageService
}

// NewCollageController is
func NewCollageController(collageSvc goexam.CollageService) *CollageController {
	return &CollageController{
		collageSvc,
	}
}

// Create is 添加 class
func (c *CollageController) Create(ctx echo.Context) error {
	collage := new(goexam.Collage)
	err := ctx.Bind(collage)
	collage, err = c.collageSvc.Create(ctx, collage)
	if err != nil {
		return ctx.String(http.StatusBadGateway, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}

// Delete is 删除
func (c *CollageController) Delete(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	if id == 0 || err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}

	err = c.collageSvc.Delete(ctx, uint(id))
	if err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	return ctx.NoContent(http.StatusOK)
}

// Update is
func (c *CollageController) UpdateName(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	if err != nil || id == 0 {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	collage := new(goexam.Collage)
	err = ctx.Bind(collage)

	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	collage, err = c.collageSvc.ChangeName(ctx, uint(id), collage.Name)
	if err != nil {
		return ctx.String(http.StatusBadGateway, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}

// Get is
func (c *CollageController) Get(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	if err != nil || id == 0 {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	class, err := c.collageSvc.GetByID(ctx, uint(id))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, class)
}

// GetList is
func (c *CollageController) GetList(ctx echo.Context) error {
	collageFilter := new(goexam.CollageFilter)
	err := ctx.Bind(collageFilter)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	collageList, err := c.collageSvc.GetList(ctx, collageFilter)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	return ctx.JSON(http.StatusOK, collageList)
}
