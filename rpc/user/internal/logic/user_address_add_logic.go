package logic

import (
	"context"

	"github.com/uncleyeung/yeung-user-center/rpc/user/internal/svc"
	"github.com/uncleyeung/yeung-user-center/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserAddressAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAddressAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddressAddLogic {
	return &UserAddressAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAddressAddLogic) UserAddressAdd(in *user.AddressUpdateReq) (*user.AddressResp, error) {
	// todo: add your logic here and delete this line

	return &user.AddressResp{}, nil
}
