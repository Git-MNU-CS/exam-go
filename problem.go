package goexam

import (
	"github.com/jinzhu/gorm"
)

type (
	// ProblemStatus is
	ProblemStatus string
	// ProblemType is
	ProblemType string
	// ProblemLevel is
	ProblemLevel uint8
)

var ProblemTypeList = map[ProblemType]string{
	Program:     "program",
	Choice:      "choice",
	FillBlank:   "fill_blank",
	ShortAnswer: "short_answer",
}

const (
	ProblemStatusEnable  ProblemStatus = "enable"
	ProblemStatusDisable ProblemStatus = "disable"

	// Program is 编程
	Program ProblemType = "program"
	// Choice is 选择
	Choice ProblemType = "choice"
	//FillBlank is 填空
	FillBlank ProblemType = "fill_blank"
	// ShortAnswer is 简答
	ShortAnswer ProblemType = "short_answer"

	// One is 1
	_ ProblemLevel = iota
	One
	// Two is 2
	Two
	// Three is 3
	Three
	// Four is 4
	Four
)

// Problem is
type Problem struct {
	gorm.Model
	Level    ProblemLevel  `json:"level" gorm:"level"`
	Type     ProblemType   `json:"type" gorm:"type"`
	Name     string        `json:"name" gorm:"name"`
	Describe string        `json:"describe" gorm:"describe"`
	CourseID uint          `json:"course_id" gorm:"course_id"`
	Status   ProblemStatus `json:"status" gorm:"status"`
	Course   *Course       `json:"course" gorm:"foreignkey:CourseID;association_foreignkey:ID"`
}

// ProblemFilter is
type ProblemFilter struct {
	BaseFilter
}

// LoadDefault is
func (p *ProblemFilter) LoadDefault() {
	p.BaseFilter.LoadDefault()
}

// ProblemService is
type ProblemService interface {
	GetList(filter *ProblemFilter) (problemList []*Problem, err error)
	Create(problem *Problem) (err error)
	Update(problem *Problem) (err error)
	Delete(id uint) (err error)
	Get(id uint) (problem *Problem, err error)
	GetByIds(ids []uint) (problemList []*Problem, err error)
}
