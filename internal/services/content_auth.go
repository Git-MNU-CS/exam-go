package services

import (
	"github.com/goexam"
)

// ContentAuthService is
type ContentAuthService struct {
	db *DB
}

var _ goexam.ContentAuthService = &ContentAuthService{}

// NewContentAuthService is
func NewContentAuthService(db *DB) *ContentAuthService {
	return &ContentAuthService{
		db,
	}
}

// Add is
func (c *ContentAuthService) Add(auth *goexam.ContentAuth) (err error) {
	err = c.db.Create(auth).Error
	return err
}

// Delete is
func (c *ContentAuthService) Delete(id uint) (err error) {
	auth := new(goexam.ContentAuth)
	err = c.db.Where("id = ?", id).Delete(auth).Error
	return err
}

// Update is
func (c *ContentAuthService) Update(auth *goexam.ContentAuth) (err error) {
	err = c.db.Model(auth).Updates(auth).Error
	return err
}

// Get is
func (c *ContentAuthService) Get(id uint) (auth *goexam.ContentAuth, err error) {
	auth = new(goexam.ContentAuth)
	err = c.db.Where("id = ?", id).First(auth).Error
	return auth, err
}

// GetByUID is
func (c *ContentAuthService) GetByUID(userID uint) (authes []*goexam.ContentAuth, err error) {
	authes = make([]*goexam.ContentAuth, 0)
	err = c.db.Where("user_id = ?", userID).Find(authes).Error
	return authes, err
}
