package services

import (
	"github.com/MNU/exam-go"
)

// ContentProblemService is
type ContentProblemService struct {
	db             *DB
	problemService goexam.ProblemService
}

var _ goexam.ContentProblemService = &ContentProblemService{}

// NewContentProblemService is
func NewContentProblemService(db *DB, problemService goexam.ProblemService) *ContentProblemService {
	return &ContentProblemService{
		db:             db,
		problemService: problemService,
	}
}

// Create is 添加一个为比赛添加一个题目
func (c *ContentProblemService) Create(cup *goexam.ContentProblem) error {
	_, err := c.problemService.Get(uint(cup.ProblemID))
	if err != nil {
		return err
	}

	err = c.db.Create(cup).Error
	return err
}

// Delete is
func (c *ContentProblemService) Delete(id uint) error {
	cup := new(goexam.ContentProblem)
	err := c.db.Where("id = ? ", id).Delete(cup).Error
	return err
}

// Update is
func (c *ContentProblemService) Update(cup *goexam.ContentProblem) error {
	err := c.db.Model(cup).Updates(cup).Error
	return err
}

// GetContentProblemIds is 获取所有考试的所有题目
func (c *ContentProblemService) GetContentProblemIds(contentID uint) ([]*uint, error) {
	cups := make([]*goexam.ContentProblem, 0)
	problemIds := make([]*uint, 0)
	err := c.db.Where("content_id = ?", contentID).Find(&cups).
		Pluck("problem_id", problemIds).Error
	return problemIds, err
}

// AddContentProblems is 添加考试题目
func (c *ContentProblemService) AddContentProblems(contentID uint, problemIds []*uint) error {
	// 没有判断考试是否存在，contentproblem 目前由content 引用
	// 判断题目是否存在

	// 添加考试题目
	contentProblem := new(goexam.ContentProblem)
	for i := 0; i < len(problemIds); i++ {
		contentProblem.ContentID = contentID
		contentProblem.ProblemID = *problemIds[i]
		contentProblem.ID = 0
		err := c.Create(contentProblem)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateContentProblems is 更改比赛的题目
func (c *ContentProblemService) UpdateContentProblems(contentID uint, newProblemIds []*uint) error {
	contentProblem := new(goexam.ContentProblem)
	err := c.DeleteByContentID(contentID)
	if err != nil {
		return err
	}
	for i := 0; i < len(newProblemIds); i++ {
		contentProblem.ID = 0
		contentProblem.ContentID = contentID
		contentProblem.ProblemID = *newProblemIds[i]
		err = c.Create(contentProblem)
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteByContentID 根据contentID删除
func (c *ContentProblemService) DeleteByContentID(contentID uint) error {
	err := c.db.Where("content_id = ?", contentID).Delete(&goexam.ContentProblem{}).Error
	return err
}
