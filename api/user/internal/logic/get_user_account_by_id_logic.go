package logic

import (
	"context"

	"github.com/uncleyeung/yeung-user-center/api/user/internal/svc"
	"github.com/uncleyeung/yeung-user-center/api/user/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserAccountByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserAccountByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserAccountByIdLogic {
	return GetUserAccountByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserAccountByIdLogic) GetUserAccountById(req types.GetByIdReq) (*types.UserAccountResp, error) {
	// todo: add your logic here and delete this line

	return &types.UserAccountResp{}, nil
}
