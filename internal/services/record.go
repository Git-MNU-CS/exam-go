package services

import (
	"github.com/MNU/exam-go"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

// RecordService is
type RecordService struct {
	db *DB
}

var _ goexam.RecordService = &RecordService{}

// NewRecordService is
func NewRecordService(db *DB) *RecordService {
	return &RecordService{
		db,
	}
}

func (r *RecordService) GetList(ctx echo.Context, filter *goexam.RecordFilter) ([]*goexam.Record, error) {
	filter.LoadDefault()

	query := r.db.DB

	if filter.UserID != 0 {
		query = query.Where("user_id = ?", filter.UserID)
	}

	if filter.ContentID != 0 {
		query = query.Where("content_id = ?", filter.ContentID)
	}

	if filter.ProblemID != 0 {
		query = query.Where("problem_id = ?", filter.ProblemID)
	}

	if filter.Page != 0 {
		query = query.Offset(filter.Page * filter.Limit)
	}

	if filter.Limit != 0 {
		query = query.Limit(filter.Limit)
	}

	records := make([]*goexam.Record, 0)

	err := query.Find(&records).Error

	if err != nil {
		return nil, err
	}

	return records, nil

}

func (r *RecordService) Get(ctx echo.Context, ID uint) (*goexam.Record, error) {
	record := new(goexam.Record)

	err := r.db.Where("id = ?", ID).First(record).Error

	if err != nil {
		return nil, errors.New("get err")
	}

	return record, nil
}

func (r *RecordService) Create(ctx echo.Context, record *goexam.Record) error {
	if record.UserID == 0 {
		return errors.New("user_id must be required")
	}

	if record.ContentID == 0 {
		return errors.New("content_id must be required")
	}

	if record.ProblemID == 0 {
		return errors.New("problem_id must be required")
	}

	if record.Result == "" {
		return errors.New("result must be required")
	}

	err := r.db.Create(record).Error

	if err != nil {
		return errors.Wrap(err, "create err")
	}

	return nil

}

func (r *RecordService) ChangeResult(ctx echo.Context, ID uint, result string) error {
	if ID == 0 {
		return errors.New("id must be required")
	}

	if result == "" {
		return errors.New("result must be required")
	}

	err := r.db.Model(&goexam.Record{}).Where("id = ?", ID).Update("result", result).Error

	if err != nil {
		return errors.Wrap(err, "update err")
	}

	return nil
}

func (r *RecordService) Delete(ctx echo.Context, ID uint) error {
	panic("implement me")
}
