package goexam

import (
	"time"

	"github.com/jinzhu/gorm"
)

// ContentStatus is 比赛状态
type ContentStatus string

const (
	// PrivateContent is 私有
	PrivateContent ContentStatus = "private"
	// PublicContent is 公开
	PublicContent ContentStatus = "public"
)

type (
	// Content is 比赛
	Content struct {
		gorm.Model
		Title      string        `json:"title" gorm:"title"`
		Describe   string        `json:"describe" gorm:"describe"`
		StartTime  time.Time     `json:"start_time" gorm:"start_time"`
		EndTime    time.Time     `json:"end_time" gorm:"end_time"`
		Status     ContentStatus `json:"status" gorm:"status"`
		ProblemIds []uint        `json:"problem_ids" gorm:"-"`
		Problems   []*Problem    `json:"problems" gorm:"-"`
		AuthIds    []uint        `json:"auth_ids" gorm:"-"`
		Autos      []uint        `json:"auths" gorm:"-"`
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
	Delete(id uint) (err error)
	Update(content *Content) (err error)
	Get(id uint) (content *Content, err error)
	GetList(filter *ContentFilter) (contentList []*Content, err error)
}
