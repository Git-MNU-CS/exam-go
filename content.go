package goexam

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	// PrivateContent is 私有
	PrivateContent = "private"
	// PublicContent is 公开
	PublicContent = "public"

	ContentStatusEnable  = "enable"
	ContentStatusDisable = "disable"
)

type (
	// Content is 比赛
	Content struct {
		gorm.Model
		Title          string         `json:"title" gorm:"title"`
		Describe       string         `json:"describe" gorm:"describe"`
		StartTime      time.Time      `json:"-" gorm:"start_time"`
		StartTimeStamp int64          `json:"start_time_stamp" gorm:"-"`
		EndTime        time.Time      `json:"-" gorm:"end_time"`
		EndTimeStamp   int64          `json:"end_time_stamp" gorm:"-"`
		Status         string         `json:"status" gorm:"status"`
		OpenDegree     string         `gorm:"column:open_degree" json:"open_degree"`
		ProblemIds     []uint         `json:"problem_ids" gorm:"-"`
		Problems       []*Problem     `json:"problems" gorm:"-"`
		UserIDs        []uint         `json:"user_ids" gorm:"-"`
		ContentUsers   []*ContentUser `json:"users" gorm:"-"`
	}
	// ContentFilter is
	ContentFilter struct {
		BaseFilter
	}
)

func (c *ContentFilter) loadDefault() {
	c.BaseFilter.LoadDefault()
}

// ContentService is
type ContentService interface {
	Create(content *Content) (*Content, error)
	Delete(id uint) error
	Update(content *Content) error
	Get(id uint) (*Content, error)
	GetList(filter *ContentFilter) ([]*Content, error)
}
