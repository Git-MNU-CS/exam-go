package student

import (
	"github.com/MNU/exam-go"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// UserController is return UserController
type RecordController struct {
	goexam.UserService
	goexam.ProblemService
	goexam.ContentService
	goexam.RecordService
}

// NewUserController is return UserController
func NewRecordController(userSvc goexam.UserService, problemSvc goexam.ProblemService, contentSvc goexam.ContentService, recordSvc goexam.RecordService) *RecordController {
	return &RecordController{
		userSvc,
		problemSvc,
		contentSvc,
		recordSvc,
	}
}

func (r *RecordController) ChangeResult(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	if err != nil || id == 0 {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	record := new(goexam.Record)
	err = ctx.Bind(record)

	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	err = r.RecordService.ChangeResult(ctx, uint(id), record.Result)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

func (r *RecordController) Submit(ctx echo.Context) error {
	record := new(goexam.Record)
	err := ctx.Bind(record)

	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	err = r.RecordService.Create(ctx, record)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

func (r *RecordController) Get(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	if err != nil || id == 0 {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}

	record, err := r.RecordService.Get(ctx, uint(id))

	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, record)
}
