package goexam

import (
	"github.com/jinzhu/gorm"
)

type (
	// Class is
	Collage struct {
		gorm.Model
		Name string `json:"name" gorm:"column:name"`
	}
	// ClassFilter is
	CollageFilter struct {
		BaseFilter
	}
)
