package services

import (
	"github.com/MNU/exam-go"
	"github.com/pkg/errors"
	"time"
)

// ContentService is
type ContentService struct {
	db *DB
}

var _ goexam.ContentService = &ContentService{}

// NewContentService is ...
func NewContentService(db *DB) *ContentService {
	return &ContentService{
		db: db,
	}
}

// Create is 添加比赛
func (c *ContentService) Create(content *goexam.Content) (*goexam.Content, error) {

	if content.Title == "" {
		return nil, errors.New("title was required")
	}

	if content.Describe == "" {
		return nil, errors.New("describe was required")
	}

	content.StartTime = time.Unix(content.StartTimeStamp, 0)
	content.EndTime = time.Unix(content.EndTimeStamp, 0)

	if content.StartTime.IsZero() {
		return nil, errors.New("start time was required")
	}

	if content.EndTime.IsZero() {
		return nil, errors.New("end time was required")
	}

	content.Status = goexam.ContentStatusEnable

	err := c.db.Create(content).Error
	if err != nil {
		return nil, err
	}

	err = c.AddContentProblems(content.ID, content.ProblemIds)

	if err != nil {
		return nil, errors.Wrap(err, "add content problem err")
	}

	// 添加考试的权限
	if content.OpenDegree == goexam.PrivateContent {
		err = c.AddContentUsers(content.ID, content.UserIDs)
	}

	if err != nil {
		return nil, errors.Wrap(err, "add content user err")
	}
	return content, nil
}

// Delete is 删除比赛
func (c *ContentService) Delete(ID uint) error {
	content := new(goexam.Content)
	err := c.db.Where("id = ?", ID).Delete(content).Error
	return err
}

// Update is update
func (c *ContentService) Update(content *goexam.Content) error {
	err := c.db.Debug().Model(content).Where("id = ?", content.ID).Updates(content).Error

	if err != nil {
		return errors.Wrap(err, "update content err")
	}
	err = c.UpdateContentProblems(content.ID, content.ProblemIds)

	if err != nil {
		return errors.Wrap(err, "update content problem err")
	}

	err = c.UpdateContentUsers(content.ID, content.UserIDs)

	if err != nil {
		return errors.Wrap(err, "update content users err")
	}

	return nil
}

// Get is 获取比赛详情
func (c *ContentService) Get(id uint) (*goexam.Content, error) {
	content := new(goexam.Content)
	err := c.db.First(content, id).Error
	if err != nil {
		return nil, err
	}
	contentProblemIds := make([]uint, 0)
	err = c.db.Debug().Where("content_id = ?", id).Model(&goexam.ContentProblem{}).Pluck("problem_id", &contentProblemIds).Error
	if err != nil {
		return nil, err
	}

	userIDs, err := c.GetContentUserIds(id)
	if err != nil {
		return nil, errors.Wrap(err, "get content user id err")
	}
	content.UserIDs = userIDs

	problemIDs, err := c.GetContentProblemIds(id)
	if err != nil {
		return nil, errors.Wrap(err, "get content problem id err")
	}

	content.ProblemIds = problemIDs

	return content, nil
}

// GetList is getlist
func (c *ContentService) GetList(filter *goexam.ContentFilter) ([]*goexam.Content, error) {
	contentList := make([]*goexam.Content, 0)

	filter.LoadDefault()

	query := c.db.LogMode(true)
	if filter.Page != 0 {
		query = query.Offset(filter.Page * filter.Limit)
	}

	if filter.Limit != 0 {
		query = query.Limit(filter.Limit)
	}

	err := query.Find(&contentList).Error
	return contentList, err
}

func (c *ContentService) UpdateContentProblems(ID uint, problemIDs []uint) error {
	err := c.DeleteContentProblems(ID)

	if err != nil {
		return err
	}

	err = c.AddContentProblems(ID, problemIDs)

	return err
}

func (c *ContentService) AddContentProblems(ID uint, problemIDs []uint) error {
	// 添加考试题目
	contentProblem := &goexam.ContentProblem{
		ContentID: ID,
	}
	for i := 0; i < len(problemIDs); i++ {
		contentProblem.ProblemID = problemIDs[i]
		contentProblem.ID = 0
		err := c.db.Create(contentProblem).Error

		if err != nil {
			return err
		}
	}

	return nil
}

func (c *ContentService) DeleteContentProblems(ID uint) error {
	err := c.db.Where("id = ?", ID).Delete(&goexam.ContentProblem{}).Error

	return err
}

func (c *ContentService) UpdateContentUsers(ID uint, userIDs []uint) error {
	err := c.DeleteContentUsers(ID)

	if err != nil {
		return err
	}

	return c.AddContentUsers(ID, userIDs)
}

func (c *ContentService) AddContentUsers(ID uint, userIDs []uint) error {
	// 添加考试权限
	contentUser := &goexam.ContentUser{
		ContentID: ID,
	}
	for i := 0; i < len(userIDs); i++ {
		contentUser.UserID = userIDs[i]
		contentUser.ID = 0
		err := c.db.Create(contentUser).Error

		if err != nil {
			return err
		}
	}

	return nil
}

func (c *ContentService) DeleteContentUsers(ID uint) error {
	err := c.db.Where("id = ?", ID).Delete(&goexam.ContentUser{}).Error

	return err
}

func (c *ContentService) GetContentProblemIds(ID uint) ([]uint, error) {
	problemsIDs := make([]uint, 0)
	contentProblems := make([]*goexam.ContentProblem, 0)
	err := c.db.Where("content_id = ?", ID).Find(&contentProblems).Pluck("problem_id", &problemsIDs).Error

	if err != nil {
		return nil, err
	}

	return problemsIDs, nil
}

func (c *ContentService) GetContentUserIds(ID uint) ([]uint, error) {
	userIDs := make([]uint, 0)

	contentUsers := make([]*goexam.ContentUser, 0)

	err := c.db.Where("content_id = ?", ID).Find(&contentUsers).Pluck("user_id", &userIDs).Error

	if err != nil {
		return nil, err
	}

	return userIDs, nil
}
