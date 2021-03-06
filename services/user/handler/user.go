package handler

import (
	"context"
	"github.com/chyidl/begin-go-micro/services/user/domain/model"
	"github.com/chyidl/begin-go-micro/services/user/domain/service"
	"github.com/chyidl/begin-go-micro/services/user/proto/user"
)

// User 绑定接口类型和实例
type User struct{
	UserDataService service.IUserDataService
}

// Register 注册
func (u *User)Register(_ context.Context, request *user.UserRegisterRequest, response *user.UserRegisterResponse) error {
	userRegister := &model.User{
		UserName: request.UserName,
		FirstName: request.FirstName,
		HashPassword: request.Pwd,
	}

	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}
	response.Message = "注册成功"
	return nil
}

// Login 登陆
func (u *User)Login(_ context.Context, request *user.UserLoginRequest, response *user.UserLoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(request.UserName, request.Pwd)
	if err != nil {
		return err
	}
	response.IsSuccess = isOk
	return nil
}

// GetUserInfo 查询
func (u *User)GetUserInfo(_ context.Context, request *user.UserInfoRequest, response *user.UserInfoResponse) error {
	userInfo, err := u.UserDataService.FindUserByName(request.UserName)
	if err != nil {
		return err
	}
	response = UserForResponse(userInfo)
	return nil
}

// UserForResponse codec
func UserForResponse(userModel *model.User) *user.UserInfoResponse {
	response := &user.UserInfoResponse{}
	response.UserId = userModel.ID
	response.UserName = userModel.UserName
	return response
}

