package controllers

import (
	"net/http"
	"strconv"

	"github.com/goexam/internal/models"

	"github.com/goexam"
	"github.com/labstack/echo"
)

// ProblemController is
type ProblemController struct {
	problemSvc goexam.ProblemService
}

// NewProblemController is
func NewProblemController(problemSvc goexam.ProblemService) *ProblemController {
	return &ProblemController{
		problemSvc,
	}
}

// Create is
func (p *ProblemController) Create(ctx echo.Context) error {
	problem := new(goexam.Problem)
	err := ctx.Bind(problem)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "参数错误")
	}
	err = p.problemSvc.Create(problem)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}

// Update is
func (p *ProblemController) Update(ctx echo.Context) error {
	problem := new(goexam.Problem)
	err := ctx.Bind(problem)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "参数错误")
	}
	err = p.problemSvc.Update(problem)
	return ctx.JSON(http.StatusOK, "成功")
}

// Delele is
func (p *ProblemController) Delele(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	err = p.problemSvc.Delete(uint(id))
	if err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	return ctx.JSON(http.StatusOK, "成功")
}

// Get is ...
func (p *ProblemController) Get(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	problem, err := p.problemSvc.Get(uint(id))
	if err != nil {
		return ctx.String(http.StatusBadRequest, "sevide")
	}
	problemResponse := new(models.ProblemResponse)
	problemResponse.BuildResponse(problem)
	return ctx.JSON(http.StatusOK, problemResponse)
}

// GetList is
func (p *ProblemController) GetList(ctx echo.Context) error {
	filter := new(goexam.ProblemFilter)
	err := ctx.Bind(filter)
	filter.LoadDefault()
	if err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	problemList, err := p.problemSvc.GetList(filter)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "服务错误")
	}
	resProblems := make([]models.ProblemResponse, len(problemList))
	for i := 0; i < len(problemList); i++ {
		resProblems[i].BuildResponse(problemList[i])
	}
	return ctx.JSON(http.StatusOK, resProblems)
}
