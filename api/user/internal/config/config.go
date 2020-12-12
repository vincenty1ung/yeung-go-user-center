package config

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Redis                    redis.RedisConf
	UserServiceRpcClientConf zrpc.RpcClientConf
	BloomRulesConf           BloomRulesConf
}
type BloomRulesConf struct {
	Rules string
}
