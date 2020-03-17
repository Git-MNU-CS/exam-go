package services

import (
	"github.com/MNU/exam-go"
	"github.com/pkg/errors"
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
func (c *ClassService) Create(class *goexam.Class) error {
	if class.Name == "" {
		return errors.New("name must required")
	}

	if class.Level == 0 {
		return errors.New("level must required")
	}

	if class.CollageID == 0 {
		return errors.New("collage id must required")
	}

	err := c.db.Create(class).Error
	return err
}

// Delete is
func (c *ClassService) Delete(id uint) error {
	err := c.db.Where("id = ?", id).Delete(&goexam.Class{}).Error
	return err
}

// Update is
func (c *ClassService) Update(class *goexam.Class) error {
	err := c.db.Model(class).Updates(class).Error
	return err
}

// Get is
func (c *ClassService) Get(id uint) (*goexam.Class, error) {
	class := new(goexam.Class)
	err := c.db.Where("id = ?", id).First(class).Error
	return class, err
}

// GetList is
func (c *ClassService) GetList(classFilter *goexam.ClassFilter) ([]*goexam.Class, error) {
	classes := make([]*goexam.Class, 0)

	query := c.db.DB
	classFilter.LoadDefault()
	if classFilter.Page != 0 {
		query.Offset(classFilter.Page * classFilter.Limit)
	}

	err := query.Limit(classFilter.Limit).Find(&classes).Error
	return classes, err
}
