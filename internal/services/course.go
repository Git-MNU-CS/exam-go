package services

import (
	"github.com/MNU/exam-go"
)

// CourseService is
type CourseService struct {
	db *DB
}

var _ goexam.CourseService = &CourseService{}

// NewCourseService is
func NewCourseService(db *DB) *CourseService {
	return &CourseService{
		db,
	}
}

// Create is
func (c *CourseService) Create(course *goexam.Course) (err error) {
	err = c.db.Create(course).Error
	return err
}

// Delete is
func (c *CourseService) Delete(id uint) (err error) {
	if id == 0 {
		return err
	}
	course := new(goexam.Course)
	err = c.db.Where("id = ?", id).Delete(course).Error
	return err
}

// Update is update
func (c *CourseService) Update(course *goexam.Course) (err error) {
	err = c.db.Model(&goexam.Course{}).Updates(course).Error
	return err
}

// Get is
func (c *CourseService) Get(id uint) (course *goexam.Course, err error) {
	course = new(goexam.Course)
	err = c.db.Debug().First(course, id).Error
	return course, err
}

// GetList is
func (c *CourseService) GetList(filter *goexam.CourseFilter) (courses []*goexam.Course, err error) {
	courses = make([]*goexam.Course, 0)
	query := c.db.Model(&goexam.Course{})
	if filter.Page != 0 {
		query = query.Offset(filter.Page * filter.Limit)
	}
	query = query.Limit(filter.Limit)
	err = query.Find(&courses).Error
	return courses, err
}
