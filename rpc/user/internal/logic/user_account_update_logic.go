package logic

import (
	"context"

	"github.com/uncleyeung/yeung-go-user-center/rpc/user/internal/svc"
	"github.com/uncleyeung/yeung-go-user-center/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserAccountUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAccountUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAccountUpdateLogic {
	return &UserAccountUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAccountUpdateLogic) UserAccountUpdate(in *user.AccountUpdateReq) (*user.AccountUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &user.AccountUpdateResp{}, nil
}
