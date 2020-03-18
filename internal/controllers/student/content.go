package student

import (
	"github.com/MNU/exam-go"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// ContentController is
type ContentController struct {
	contentSvc goexam.ContentService
	problemSvc goexam.ProblemService
	userSvc    goexam.UserService
}

// NewContentController is
func NewContentController(contentSvc goexam.ContentService, problemSvc goexam.ProblemService, userSvc goexam.UserService) *ContentController {
	return &ContentController{
		contentSvc,
		problemSvc,
		userSvc,
	}
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

	problems, err := c.problemSvc.GetByIds(content.ProblemIds)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "get problem err")
	}

	content.Problems = problems

	return ctx.JSON(http.StatusOK, content)
}
