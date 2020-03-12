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

// ContentProblemService is
type ContentProblemService interface {
	Create(cup *ContentProblem) error
	Delete(id uint) error
	Update(cup *ContentProblem) (err error)
	GetContentProblemIds(contentID uint) ([]*uint, error)
	AddContentProblems(contentID uint, problemIds []*uint) error
	UpdateContentProblems(contentID uint, newProblemIds []*uint) error
	DeleteByContentID(contentID uint) error
}

// TableName is
func (u *ContentProblem) TableName() string {
	return "content_problems"
}
