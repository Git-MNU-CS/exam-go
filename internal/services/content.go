package services

import (
	"github.com/goexam"
)

// ContentService is
type ContentService struct {
	db                    *DB
	problemService        goexam.ProblemService
	contentProblemService goexam.ContentProblemService
}

var _ goexam.ContentService = &ContentService{}

// NewContentService is ...
func NewContentService(db *DB, problemService goexam.ProblemService, contentProblemService goexam.ContentProblemService) *ContentService {
	return &ContentService{
		db:                    db,
		problemService:        problemService,
		contentProblemService: contentProblemService,
	}
}

// Create is 添加比赛
func (c *ContentService) Create(content *goexam.Content) (*goexam.Content, error) {
	err := c.db.Create(content).Error
	if err != nil {
		return nil, err
	}
	// 添加考试题目
	err = c.contentProblemService.AddContentProblems(content.ID, content.ProblemIds)
	if err != nil {
		return nil, err
	}
	// 添加考试的权限
	if content.Status == goexam.PrivateContent {
		err = c.UpdateContentAuth(content.ID, content.AuthIds)
		if err != nil {
			return nil, err
		}
	}
	return content, err
}

// Delete is 删除比赛
func (c *ContentService) Delete(id uint) (err error) {
	content := new(goexam.Content)
	err = c.db.Where("id = ?", id).Delete(content).Error
	if err != nil {
		return err
	}
	err = c.contentProblemService.DeleteByContentID(id)
	return err
}

// Update is update
func (c *ContentService) Update(content *goexam.Content) (err error) {
	err = c.db.Debug().Model(content).Where("id = ?", content.ID).Updates(content).Error
	if err != nil {
		return err
	}
	err = c.contentProblemService.UpdateContentProblems(content.ID, content.ProblemIds)
	if err != nil {
		return err
	}
	return err
}

// Get is 获取比赛详情
func (c *ContentService) Get(id uint) (content *goexam.Content, err error) {
	content = new(goexam.Content)
	err = c.db.First(content, id).Error
	if err != nil {
		return nil, err
	}
	contentProblemIds := make([]uint, 0)
	err = c.db.Debug().Where("content_id = ?", id).Model(&goexam.ContentProblem{}).Pluck("problem_id", &contentProblemIds).Error
	if err != nil {
		return nil, err
	}
	problems := make([]*goexam.Problem, 0)
	err = c.db.Debug().Where("id in (?)", contentProblemIds).Find(&problems).Error
	if err != nil {
		return nil, err
	}
	content.Problems = problems
	return content, err
}

// GetList is getlist
func (c *ContentService) GetList(filter *goexam.ContentFilter) (contentList []*goexam.Content, err error) {
	contentList = make([]*goexam.Content, 0)
	err = c.db.Find(contentList).Error
	return contentList, err
}

// UpdateContentAuth 添加考试权限
func (c *ContentService) UpdateContentAuth(id uint, authIds []uint) (err error) {
	contentAuth := &goexam.ContentAuth{}
	for i := 0; i < len(authIds); i++ {
		contentAuth.ContentID = id
		contentAuth.UserID = authIds[i]
		contentAuth.ID = 0
		// err = c.db.Model(contentAuth).Where("content_id = ? and user_id = ?", id, authIds[i]).FirstOrCreate(contentAuth).Error
		err = c.db.Create(contentAuth).Error
		if err != nil {
			return err
		}
	}
	return nil
}
