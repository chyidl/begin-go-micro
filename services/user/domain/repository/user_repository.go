package repository

import (
	"github.com/chyidl/begin-go-micro/services/user/domain/model"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	// InitTable 初始化数据表
	InitTable() error
	// FindUserByName 根据用户名称查询用户信息
	FindUserByName(string) (*model.User, error)
	// FindUserByID 根据用户ID查询用户信息
	FindUserByID(int64) (*model.User, error)
	// CreateUser 创建用户
	CreateUser(*model.User) (int64,error)
	// DeleteUserByID 根据用户ID删除用户
	DeleteUserByID(int64) error
	// UpdateUser 更新用户信息
	UpdateUser(*model.User) error
	// FindAll 查找所有用户
	FindAll() ([]model.User, error)
}

// NewUserRepository 创建UserRepository
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb: db}
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

// InitTable 初始化表
func (u *UserRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.User{}).Error
}

// FindUserByName 根据用户名称查找用户信息
func (u *UserRepository) FindUserByName(name string) (user *model.User, err error) {
	user  = &model.User{}
	return user, u.mysqlDb.Where("user_name = ?", name).Find(user).Error
}

// FindUserByID 根据用户ID查找用户信息
func (u *UserRepository) FindUserByID(userID int64) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.First(user, userID).Error
}

// CreateUser 创建用户
func (u *UserRepository) CreateUser(user *model.User) (userID int64, err error) {
	return user.ID, u.mysqlDb.Create(user).Error
}

// DeleteUserByID 根据用户ID删除用户
func (u *UserRepository) DeleteUserByID(userID int64) error {
	return u.mysqlDb.Where("id = ?", userID).Delete(&model.User{}).Error
}

// UpdateUser 更新用户信息
func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDb.Model(user).Update(&user).Error
}

// FindAll 查找所有用户
func (u *UserRepository) FindAll() (userAll []model.User, err error) {
	return userAll, u.mysqlDb.Find(&userAll).Error
}