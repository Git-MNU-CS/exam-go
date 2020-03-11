package goexam

import (
	"github.com/jinzhu/gorm"
)

type (
	// Course is
	Course struct {
		gorm.Model
		Name string `json:"name" gorm:"column:name"`
	}

	// CourseFilter is
	CourseFilter struct {
		BaseFilter
	}
)

func (c *CourseFilter) loadDefault() {
	c.BaseFilter.LoadDefault()
}

// CourseService is
type CourseService interface {
	Create(course *Course) (err error)
	Delete(id uint) (err error)
	Update(course *Course) (err error)
	Get(id uint) (course *Course, err error)
	GetList(filter *CourseFilter) (courses []*Course, err error)
}

// TableName is
func (u *Course) TableName() string {
	return "courses"
}
