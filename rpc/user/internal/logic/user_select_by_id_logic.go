package logic

import (
	"context"

	"github.com/uncleyeung/yeung-user-center/rpc/user/internal/svc"
	"github.com/uncleyeung/yeung-user-center/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserSelectByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserSelectByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSelectByIdLogic {
	return &UserSelectByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserSelectByIdLogic) UserSelectById(in *user.GetByIdReq) (*user.UserGetByIdResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserGetByIdResp{}, nil
}
