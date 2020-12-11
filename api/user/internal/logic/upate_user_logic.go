package logic

import (
	"context"

	"github.com/uncleyeung/yeung-user-center/api/user/internal/svc"
	"github.com/uncleyeung/yeung-user-center/api/user/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpateUserLogic {
	return UpateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpateUserLogic) UpateUser(req types.UserUpdateRequest) (*types.DmlResponse, error) {
	// todo: add your logic here and delete this line

	return &types.DmlResponse{}, nil
}
