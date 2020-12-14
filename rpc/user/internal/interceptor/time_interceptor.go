package interceptor

import (
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"google.golang.org/grpc"
	"time"
)

func TimeInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	st := time.Now()
	resp, err = handler(ctx, req)
	logx.Info(fmt.Sprintf("TimeInterceptor==method: %s time: %v\n", info.FullMethod, time.Since(st)))
	return resp, err
}
