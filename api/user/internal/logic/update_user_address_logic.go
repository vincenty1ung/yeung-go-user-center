package logic

import (
	"context"

	"github.com/uncleyeung/yeung-user-center/api/user/internal/svc"
	"github.com/uncleyeung/yeung-user-center/api/user/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateUserAddressLogic {
	return UpdateUserAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserAddressLogic) UpdateUserAddress(req types.UserAddressUpdateReq) error {
	// todo: add your logic here and delete this line

	return nil
}
