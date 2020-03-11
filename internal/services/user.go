package services

import (
	"github.com/goexam"
)

// UserService is
type UserService struct {
	db       *DB
	classSvc goexam.ClassService
}

var _ goexam.UserService = &UserService{}

// NewUserService is return Userservice *
func NewUserService(db *DB, classSvc goexam.ClassService) *UserService {
	return &UserService{
		db:       db,
		classSvc: classSvc,
	}
}

// Login is 用户登陆
func (u *UserService) Login(username string, password string) (err error) {
	user := new(goexam.User)
	err = u.db.Where("username = ? and passwd = ?", username, password).Find(user).Error
	return err
}

// Create is 添加用户
func (u *UserService) Create(user *goexam.User) (err error) {
	_, err = u.classSvc.Get(user.ClassID)
	if err != nil {
		return err
	}
	err = u.db.Create(user).Error
	return err
}

// Delete is 删除用户
func (u *UserService) Delete(id uint) (err error) {
	user := new(goexam.User)
	err = u.db.Where("id = ?", id).Delete(user).Error
	return err
}

// Update is 更改用户
func (u *UserService) Update(user *goexam.User) (err error) {
	classID := user.ClassID
	if classID != 0 {
		_, err = u.classSvc.Get(classID)
		if err != nil {
			return err
		}
	}
	err = u.db.Model(user).Updates(user).Error
	return err
}

// Get is 获取用户信息
func (u *UserService) Get(id uint) (user *goexam.User, err error) {
	user = new(goexam.User)
	err = u.db.Preload("Class").First(user, id).Error
	if err != nil {
		return user, err
	}
	return user, err
}

// GetList is 获取用户列表
func (u *UserService) GetList(userFilter *goexam.UserFilter) (userList []*goexam.User, err error) {
	userList = make([]*goexam.User, 0)
	query := u.db.Model(&goexam.User{}).Preload("Class")
	if userFilter.Page != 0 {
		query = query.Offset(userFilter.Page * userFilter.Limit)
	}
	query = query.Limit(userFilter.Limit)
	err = query.Find(&userList).Error
	return userList, err
}
