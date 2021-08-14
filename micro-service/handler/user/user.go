package user

import (
	"context"
	"micro-service/internal/service"
	protoUser "micro-service/proto/user"
)

type User struct{}

func (u *User) Login(ctx context.Context, req *protoUser.LoginRequest, rsp *protoUser.LoginResponse) error {
	if req.Email != "gopher@88.com" || req.Password != "123456" {
		rsp.Username = "Sorry " + req.Email
		return nil
	}
	rsp.Username = "Welcome " + req.Email

	return nil
}

func (u *User) Register(ctx context.Context, req *protoUser.User, rsp *protoUser.RegisterResponse) error {
	registerReq := &protoUser.User{
		Username: req.Username,
		Email: req.Email,
		Password: req.Password,
	}
	svc := service.New()
	affected, err := svc.Add(registerReq)
	if err != nil {
		return err
	}
	rsp.Code = 10000
	rsp.Message = "success"
	rsp.Affected = affected
	return nil
}

func (u *User) UserInfo(ctx context.Context, req *protoUser.User, rsp *protoUser.UserInfoResponse) error {
	registerReq := &protoUser.User{Id: req.Id}
	svc := service.New()
	data, err := svc.Get(registerReq)
	if err != nil {
		return err
	}
	rsp.Code = 10000
	rsp.Message = "success"
	rsp.Data = data
	return nil
}
