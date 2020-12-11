package logic

import (
	"context"

	"github.com/uncleyeung/yeung-user-center/api/user/internal/svc"
	"github.com/uncleyeung/yeung-user-center/api/user/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateUserAccountLogic {
	return UpdateUserAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserAccountLogic) UpdateUserAccount(req types.UserAccountUpdateReq) error {
	// todo: add your logic here and delete this line

	return nil
}
