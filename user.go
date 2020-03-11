package goexam

import (
	"github.com/jinzhu/gorm"
)

// Role is
type Role string

const (
	// Student is init
	StudentRole Role = "student"
	// Teacher is
	TeacherRole Role = "teacher"

	AdminRole Role = "admin"
)

type (
	// User is
	User struct {
		gorm.Model
		Role     Role   `json:"role" gorm:"column:role"`
		Username string `json:"username" gorm:"column:username"`
		Password string `json:"password" gorm:"column:password"`
		Name     string `json:"name" gorm:"column:name"`
		ClassID  uint   `json:"class_id" gorm:"column:class_id"`
		Class    *Class `json:"class" gorm:"foreignkey:ClassID;association_foreignkey:ID"`
	}

	// UserFilter is
	UserFilter struct {
		BaseFilter
	}
)

func (u *UserFilter) loadDefault() {
	u.BaseFilter.LoadDefault()
}

// UserService is
type UserService interface {
	Login(username string, password string) (err error)
	Create(user *User) (err error)
	Delete(id uint) (err error)
	Update(user *User) (err error)
	Get(id uint) (user *User, err error)
	GetList(userFilter *UserFilter) (userList []*User, err error)
}

// TableName is
func (u *User) TableName() string {
	return "users"
}
