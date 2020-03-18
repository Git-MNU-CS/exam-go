package goexam

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// Record is
type Record struct {
	gorm.Model
	UserID    uint   `json:"user_id" gorm:"column:user_id"`
	ContentID uint   `json:"content_id" gorm:"column:content_id"`
	ProblemID uint   `json:"problem_id" gorm:"column:problem_id"`
	Result    string `json:"result" gorm:"column:result"`
}

type RecordFilter struct {
	BaseFilter
	UserID    uint `json:"user_id"`
	ContentID uint `json:"content_id"`
	ProblemID uint `json:"problem_id"`
}

// RecordService is
type RecordService interface {
	GetList(ctx echo.Context, filter *RecordFilter) ([]*Record, error)
	Get(ctx echo.Context, ID uint) (*Record, error)
	Create(ctx echo.Context, record *Record) error
	ChangeResult(ctx echo.Context, ID uint, result string) error
	Delete(ctx echo.Context, ID uint) error
}

// TableName is
func (u *Record) TableName() string {
	return "records"
}
