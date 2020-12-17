package svc

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/uncleyeung/yeung-go-user-center/api/user/internal/bloom"
	"github.com/uncleyeung/yeung-go-user-center/api/user/internal/config"
	"github.com/uncleyeung/yeung-go-user-center/api/user/internal/middleware"
	"github.com/uncleyeung/yeung-go-user-center/rpc/user/userservice"
	"strings"
)

type ServiceContext struct {
	Config      config.Config
	UserService userservice.UserService
	BloomCheck  rest.Middleware
	UserCheck   rest.Middleware
	LoginCheck  rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	bloomerUser := bloom.MustNewBloom(redis.NewRedis(c.Redis.Host, c.Redis.Type, c.Redis.Pass))
	split := strings.Split(c.BloomRulesConf.Rules, ",")
	err := bloomerUser.AddBloomRules(split...)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:      c,
		UserService: userservice.NewUserService(zrpc.MustNewClient(c.UserServiceRpcClientConf)),
		BloomCheck:  middleware.NewBloomCheckMiddleware(bloomerUser).Handle,
		UserCheck:   middleware.NewUserCheckMiddleware().Handle,
		LoginCheck:  middleware.NewLoginCheckMiddleware().Handle,
	}
}
