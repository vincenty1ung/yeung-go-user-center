package logic

import (
	"context"

	"github.com/uncleyeung/yeung-user-center/api/user/internal/svc"
	"github.com/uncleyeung/yeung-user-center/api/user/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ListUserAddressByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListUserAddressByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListUserAddressByUserIdLogic {
	return ListUserAddressByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUserAddressByUserIdLogic) ListUserAddressByUserId(req types.GetByUserIdReq) (*types.AddressListResp, error) {
	// todo: add your logic here and delete this line

	return &types.AddressListResp{}, nil
}
