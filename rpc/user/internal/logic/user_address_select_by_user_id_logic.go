package logic

import (
	"context"

	"github.com/uncleyeung/yeung-user-center/rpc/user/internal/svc"
	"github.com/uncleyeung/yeung-user-center/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserAddressSelectByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAddressSelectByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddressSelectByUserIdLogic {
	return &UserAddressSelectByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAddressSelectByUserIdLogic) UserAddressSelectByUserId(in *user.GetByIdReq) (*user.AddressListResp, error) {
	// todo: add your logic here and delete this line

	return &user.AddressListResp{}, nil
}
