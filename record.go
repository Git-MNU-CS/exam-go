package goexam

import (
	"github.com/jinzhu/gorm"
)

// Record is
type Record struct {
	gorm.Model
	UserID    uint   `json:"user_id" gorm:"column:user_id"`
	ContentID uint   `json:"content_id" gorm:"column:content_id"`
	ProblemID uint   `json:"problem_id" gorm:"column:problem_id"`
	Result    string `json:"result" gorm:"column:result"`
}

// RecordService is
type RecordService interface {
}

// TableName is
func (u *Record) TableName() string {
	return "records"
}
