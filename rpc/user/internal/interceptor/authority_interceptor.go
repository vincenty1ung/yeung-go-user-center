package interceptor

import (
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"google.golang.org/grpc"
)

//权限拦截器
func AuthorityInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	logx.Info(fmt.Sprintf("AuthorityInterceptor==权限校验通过info: %+v\n", info))
	return handler(ctx, req)
}
