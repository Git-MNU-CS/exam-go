package models

import (
	"github.com/goexam"
)

type (
	// UserLoginRequest is
	UserLoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// UserResponse is
	UserResponse struct {
		ID        uint        `json:"id"`
		RoleID    goexam.Role `json:"role_id"`
		Username  string      `json:"username"`
		Name      string      `json:"name"`
		ClassID   uint        `json:"class_id"`
		ClassName string      `json:"class_name"`
		CreatedAt string      `json:"created_at"`
		UpdatedAt string      `json:"updated_at"`
	}
)

// BuildResponse is
func (u *UserResponse) BuildResponse(user *goexam.User) {
	*u = UserResponse{
		ID:        user.ID,
		RoleID:    user.RoleID,
		Username:  user.Username,
		Name:      user.Name,
		ClassID:   user.ClassID,
		ClassName: user.Class.Name,
		CreatedAt: user.CreatedAt.Format(TimeFormat),
		UpdatedAt: user.UpdatedAt.Format(TimeFormat),
	}
}
