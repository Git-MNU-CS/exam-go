package goexam

import (
	"github.com/jinzhu/gorm"
)

// ContentAuth is
type ContentAuth struct {
	gorm.Model
	ContentID uint `json:"content_id" gorm:"column:content_id"`
	UserID    uint `json:"user_id" gorm:"column:user_id"`
}

// ContentAuthService is
type ContentAuthService interface {
}

// TableName is
func (u *ContentAuth) TableName() string {
	return "content_authorities"
}
