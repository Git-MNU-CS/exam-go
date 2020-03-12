package services

import "github.com/MNU/exam-go"

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

// Create is
func (r *RecordService) Create(record goexam.Record) (err error) {
	return err
}

// Delete is
func (r *RecordService) Delete() (err error) {
	return err
}

// Update is
func (r *RecordService) Update() (err error) {
	return err
}

// GetByID is
func (r *RecordService) GetByID() (err error) {
	return err
}

// GetByUID is
func (r *RecordService) GetByUID() (err error) {
	return err
}

// GetList is
func (r *RecordService) GetList() (err error) {
	return err
}
