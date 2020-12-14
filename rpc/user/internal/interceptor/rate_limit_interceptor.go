package interceptor

import (
	"context"
	"errors"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

//限流算法:https://www.cnblogs.com/gnivor/p/10623028.html
var limiter = rate.NewLimiter(rate.Limit(1), 5)

func RateLimitInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	logx.Info(fmt.Sprintf("RateLimitInterceptor==method: %s", info.FullMethod))
	if !limiter.Allow() {
		logx.Info(fmt.Sprintf("RateLimitInterceptor==method: %s ===========触发限流", info.FullMethod))
		return nil, errors.New("触发限流策略")
	}
	return handler(ctx, req)
}
