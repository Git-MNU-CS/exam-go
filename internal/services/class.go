package services

import (
	"github.com/goexam"
)

// ClassService is
type ClassService struct {
	db *DB
}

var _ goexam.ClassService = &ClassService{}

// NewClassService is
func NewClassService(db *DB) *ClassService {
	return &ClassService{
		db,
	}
}

// Create is
func (c *ClassService) Create(class *goexam.Class) (err error) {
	err = c.db.Create(class).Error
	return err
}

// Delete is
func (c *ClassService) Delete(id uint) (err error) {
	class := new(goexam.Class)
	err = c.db.Where("id = ?", id).Delete(class).Error
	return err
}

// Update is
func (c *ClassService) Update(class *goexam.Class) (err error) {
	err = c.db.Model(class).Updates(class).Error
	return err
}

// Get is
func (c *ClassService) Get(id uint) (class *goexam.Class, err error) {
	class = new(goexam.Class)
	err = c.db.Where("id = ?", id).First(class).Error
	return class, err
}

// GetList is
func (c *ClassService) GetList(classFilter *goexam.ClassFilter) (classes []*goexam.Class, err error) {
	classes = make([]*goexam.Class, 0)
	err = c.db.Find(&classes).Limit(classFilter.Limit).Error
	return classes, err
}
