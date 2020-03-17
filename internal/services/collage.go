package services

import (
	"github.com/MNU/exam-go"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type CollageService struct {
	db *DB
}

var _ goexam.CollageService = &CollageService{}

func NewCollageService(db *DB) *CollageService {
	return &CollageService{
		db: db,
	}
}

func (c *CollageService) Create(ctx echo.Context, collage *goexam.Collage) (*goexam.Collage, error) {

	if collage.Name == "" {
		return nil, errors.New("collage name must require")
	}
	err := c.db.Create(collage).Error

	return collage, err
}

func (c *CollageService) GetByID(ctx echo.Context, ID uint) (*goexam.Collage, error) {

	collage := new(goexam.Collage)

	err := c.db.Where("id = ?", ID).First(collage).Error

	return collage, err
}

func (c *CollageService) ChangeName(ctx echo.Context, ID uint, name string) (*goexam.Collage, error) {
	err := c.db.Model(goexam.Collage{}).Where("id = ?", ID).Update("name", name).Error

	return nil, err
}

func (c *CollageService) GetList(ctx echo.Context, filter *goexam.CollageFilter) ([]*goexam.Collage, error) {
	list := make([]*goexam.Collage, 0)

	filter.LoadDefault()
	query := c.db.DB

	if filter.Name != "" {
		query = query.Where("name like %?%", filter.Name)
	}

	if filter.Limit != 0 {
		query = query.Limit(filter.Limit)
	}

	err := query.Find(&list).Error
	return list, err
}

func (c *CollageService) Delete(ctx echo.Context, ID uint) error {
	err := c.db.Where("id = ?", ID).Delete(&goexam.Collage{}).Error

	return err
}
