// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for User service

func NewUserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for User service

type UserService interface {
	// 注册
	Register(ctx context.Context, in *UserRegisterRequest, opts ...client.CallOption) (*UserRegisterResponse, error)
	// 登陆
	Login(ctx context.Context, in *UserLoginRequest, opts ...client.CallOption) (*UserLoginResponse, error)
	// 查询
	GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...client.CallOption) (*UserInfoResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Register(ctx context.Context, in *UserRegisterRequest, opts ...client.CallOption) (*UserRegisterResponse, error) {
	req := c.c.NewRequest(c.name, "User.Register", in)
	out := new(UserRegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Login(ctx context.Context, in *UserLoginRequest, opts ...client.CallOption) (*UserLoginResponse, error) {
	req := c.c.NewRequest(c.name, "User.Login", in)
	out := new(UserLoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...client.CallOption) (*UserInfoResponse, error) {
	req := c.c.NewRequest(c.name, "User.GetUserInfo", in)
	out := new(UserInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	// 注册
	Register(context.Context, *UserRegisterRequest, *UserRegisterResponse) error
	// 登陆
	Login(context.Context, *UserLoginRequest, *UserLoginResponse) error
	// 查询
	GetUserInfo(context.Context, *UserInfoRequest, *UserInfoResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Register(ctx context.Context, in *UserRegisterRequest, out *UserRegisterResponse) error
		Login(ctx context.Context, in *UserLoginRequest, out *UserLoginResponse) error
		GetUserInfo(ctx context.Context, in *UserInfoRequest, out *UserInfoResponse) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Register(ctx context.Context, in *UserRegisterRequest, out *UserRegisterResponse) error {
	return h.UserHandler.Register(ctx, in, out)
}

func (h *userHandler) Login(ctx context.Context, in *UserLoginRequest, out *UserLoginResponse) error {
	return h.UserHandler.Login(ctx, in, out)
}

func (h *userHandler) GetUserInfo(ctx context.Context, in *UserInfoRequest, out *UserInfoResponse) error {
	return h.UserHandler.GetUserInfo(ctx, in, out)
}
