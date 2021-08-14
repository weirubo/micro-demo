package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	protoUser "micro-client/proto/user"
	"net/http"
)

type User struct{}

func (u *User) Login(ctx *gin.Context) {
	param := new(protoUser.User)
	err := ctx.ShouldBind(param)
	if err != nil {
		log.Debug(err)
		return
	}

	client := NewUserClient()

	// rpc 调用远程服务的方法
	resp, err := client.Login(context.TODO(), &protoUser.LoginRequest{Email: param.Email, Password: param.Password})
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": resp.Username,
	})

}

func (u *User) Register(ctx *gin.Context) {
	param := new(protoUser.User)
	err := ctx.ShouldBind(param)
	if err != nil {
		log.Debug(err)
		return
	}

	client := NewUserClient()

	resp, err := client.Register(context.TODO(), &protoUser.User{Username: param.Username, Email: param.Email, Password: param.Password})
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"message": resp.Message,
		"data":    resp.Affected,
	})
}

func (u *User) UserInfo(ctx *gin.Context) {
	param := new(protoUser.User)
	err := ctx.ShouldBind(param)
	if err != nil {
		log.Debug(err)
		return
	}

	client := NewUserClient()

	resp, err := client.UserInfo(context.TODO(), &protoUser.User{Id: param.Id})
	if err != nil {
		fmt.Printf("client.Get() error:%v\n", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"message": resp.Message,
		"data":    resp.Data,
	})
}

func NewUserClient() protoUser.UserService {
	// 创建服务
	service := grpc.NewService(
		micro.Name("go.micro.cli.demo"),
		micro.Version("v0.0.0"),
	)

	// 创建客户端
	userClient := protoUser.NewUserService("go.micro.srv.demo", service.Client())

	return userClient
}
