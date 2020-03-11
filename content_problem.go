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
	Create(cup *ContentProblem) (err error)
	Delete(id uint) (err error)
	Update(cup *ContentProblem) (err error)
	GetContentProblemIds(contentID uint) (problemIds []uint, err error)
	AddContentProblems(contentID uint, problemIds []uint) (err error)
	UpdateContentProblems(contentID uint, newProblemIds []uint) (err error)
	DeleteByContentID(contentID uint) (err error)
}

// TableName is
func (u *ContentProblem) TableName() string {
	return "content_problems"
}
