package service

import (
	"errors"
	"github.com/chyidl/begin-go-micro/services/user/domain/model"
	"github.com/chyidl/begin-go-micro/services/user/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserDataService interface {
	AddUser(*model.User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(user *model.User, isChangePwd bool) (err error)
	FindUserByName(string) (*model.User, error)
	CheckPwd(userName string, password string) (isOk bool, err error)
}

func NewUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository: userRepository}
}


type UserDataService struct {
	UserRepository repository.IUserRepository
}

// GeneratePassword 加密用户密码
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

// ValidatePassword 验证用户密码
func ValidatePassword(userPassword string, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("password is increate")
	}
	return true, nil
}

// AddUser 插入用户
func (u *UserDataService) AddUser(user *model.User) (userID int64, err error) {
	// 对用户密码加密
	pwdByte, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return user.ID, err
	}

	user.HashPassword = string(pwdByte)
	return u.UserRepository.CreateUser(user)
}

// DeleteUser 删除用户
func (u *UserDataService) DeleteUser(userID int64) error {
	return u.UserRepository.DeleteUserByID(userID)
}

// UpdateUser 更新用户
func (u *UserDataService)UpdateUser(user *model.User, isChangePwd bool) (err error) {
	// 判断是否更新密码
	if isChangePwd {
		pwdByte, err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdByte)
	}
	// 更新日志
	return u.UserRepository.UpdateUser(user)
}

// FindUserByName 根据用户名称查找用户
func (u *UserDataService) FindUserByName(userName string) (user *model.User, err error) {
	return u.UserRepository.FindUserByName(userName)
}

// CheckPwd 验证密码
func (u *UserDataService) CheckPwd(userName string, password string) (isOk bool, err error) {
	user, err := u.UserRepository.FindUserByName(userName)
	if err != nil {
		return false, err
	}
	return ValidatePassword(password, user.HashPassword)
}