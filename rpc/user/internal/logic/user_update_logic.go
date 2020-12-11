package logic

import (
	"context"

	"github.com/uncleyeung/yeung-user-center/rpc/user/internal/svc"
	"github.com/uncleyeung/yeung-user-center/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdateLogic) UserUpdate(in *user.UserUpdateReq) (*user.UserUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserUpdateResp{}, nil
}
