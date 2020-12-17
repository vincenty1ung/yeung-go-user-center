package logic

import (
	"context"

	"github.com/uncleyeung/yeung-go-user-center/api/user/internal/svc"
	"github.com/uncleyeung/yeung-go-user-center/api/user/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserAccountByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserAccountByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserAccountByUserIdLogic {
	return GetUserAccountByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserAccountByUserIdLogic) GetUserAccountByUserId(req types.GetByUserIdReq) (*types.UserAccountResp, error) {
	// todo: add your logic here and delete this line

	return &types.UserAccountResp{}, nil
}
