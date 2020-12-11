package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/uncleyeung/yeung-user-center/api/user/internal/config"
	"github.com/uncleyeung/yeung-user-center/rpc/user/userservice"
)

type ServiceContext struct {
	Config      config.Config
	UserService userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		UserService: userservice.NewUserService(zrpc.MustNewClient(c.UserServiceRpcClientConf)),
	}
}
