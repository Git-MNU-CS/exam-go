package goexam

import (
	"github.com/jinzhu/gorm"
)

type (
	// Class is
	Class struct {
		gorm.Model
		Name      string   `json:"name" gorm:"column:name"`
		Level     uint8    `json:"level" gorm:"column:level"`
		CollageID uint     `json:"collage_id" gorm:"column:collage_id"`
		Collage   *Collage `json:"collage" gorm:"-"`
	}
	// ClassFilter is
	ClassFilter struct {
		BaseFilter
	}
)

// LoadDefault is
func (c *ClassFilter) LoadDefault() {
	c.BaseFilter.LoadDefault()
}

// ClassService is
type ClassService interface {
	Create(class *Class) (err error)
	Delete(id uint) (err error)
	Update(class *Class) (err error)
	Get(id uint) (class *Class, err error)
	GetList(classFilter *ClassFilter) (classes []*Class, err error)
}

// TableName is
func (u *Class) TableName() string {
	return "classes"
}
