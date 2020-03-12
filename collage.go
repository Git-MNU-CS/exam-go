package goexam

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	// Class is
	Collage struct {
		gorm.Model
		Name string `json:"name" gorm:"column:name"`
	}
	// ClassFilter is
	CollageFilter struct {
		Name string
		BaseFilter
	}
)

type CollageService interface {
	Create(ctx echo.Context, collage *Collage) (*Collage, error)
	GetByID(ctx echo.Context, ID uint) (*Collage, error)
	ChangeName(ctx echo.Context, ID uint, name string) (*Collage, error)
	GetList(ctx echo.Context, filter *CollageFilter) ([]*Collage, error)
}
