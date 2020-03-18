package student

import (
	"github.com/MNU/exam-go"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// UserController is return UserController
type UserController struct {
	goexam.UserService
	goexam.ClassService
}

// NewUserController is return UserController
func NewUserController(userSvc goexam.UserService, classSvc goexam.ClassService) *UserController {
	return &UserController{
		userSvc,
		classSvc,
	}
}

// Get is
func (uc *UserController) Get(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	if id == 0 || err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	user, err := uc.UserService.Get(uint(id))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "系统错误")
	}
	return ctx.JSON(http.StatusOK, user)
}

// Login is
func (uc *UserController) Login(ctx echo.Context) error {
	user := new(goexam.User)
	err := ctx.Bind(user)

	if err != nil {
		return err
	}
	err = uc.UserService.Login(user.Account, user.Password)
	return ctx.JSON(http.StatusOK, "登录成功")
}

func (uc *UserController) Logout(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}
