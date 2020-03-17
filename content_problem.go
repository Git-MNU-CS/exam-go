package goexam

import (
	"github.com/jinzhu/gorm"
)

// ContentProblem is
type ContentProblem struct {
	gorm.Model
	ContentID uint `json:"content_id" gorm:"column:content_id"`
	ProblemID uint `json:"problem_id" gorm:"column:problem_id"`
}

// TableName is
func (u *ContentProblem) TableName() string {
	return "content_problems"
}
