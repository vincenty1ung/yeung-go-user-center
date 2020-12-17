package logic

import (
	"context"

	"github.com/uncleyeung/yeung-go-user-center/rpc/user/internal/svc"
	"github.com/uncleyeung/yeung-go-user-center/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDeleteLogic) UserDelete(in *user.GetByIdReq) (*user.UserDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserDeleteResp{}, nil
}
