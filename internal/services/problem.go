package services

import (
	"github.com/MNU/exam-go"
)

// ProblemService is
type ProblemService struct {
	db *DB
}

var _ goexam.ProblemService = &ProblemService{}

// NewProblemService is return ProblemServiceInstance
func NewProblemService(db *DB) *ProblemService {
	return &ProblemService{
		db: db,
	}
}

// GetList is 题目列表
func (p *ProblemService) GetList(filter *goexam.ProblemFilter) (problemList []*goexam.Problem, err error) {
	problemList = make([]*goexam.Problem, 0)
	problem := new(goexam.Problem)
	query := p.db.Model(problem).Preload("Course")
	if filter.PrefixKey != "" {
		query = query.Where("name like ?", "%"+filter.PrefixKey)
	}
	if filter.Page != 0 {
		query = query.Offset(filter.Page * filter.Limit)
	}
	err = query.Limit(filter.Limit).Find(&problemList).Error
	return problemList, err
}

// Create is 添加题目
func (p *ProblemService) Create(problem *goexam.Problem) error {
	err := p.db.Create(problem).Error
	return err
}

// Update is 编辑题目
func (p *ProblemService) Update(problem *goexam.Problem) error {
	err := p.db.Updates(problem).Error
	return err
}

// Delete is 删除题目
func (p *ProblemService) Delete(ID uint) error {
	problem := new(goexam.Problem)
	err := p.db.Where("id = ?", ID).Delete(problem).Error
	return err
}

// Get is
func (p *ProblemService) Get(ID uint) (*goexam.Problem, error) {
	problem := new(goexam.Problem)
	err := p.db.Debug().Preload("Course").First(problem, ID).Error
	return problem, err
}

// GetByIds is
func (p *ProblemService) GetByIds(ids []uint) ([]*goexam.Problem, error) {
	problemList := make([]*goexam.Problem, 0)
	err := p.db.Where("id in (?)", ids).Find(problemList).Error
	return problemList, err
}
