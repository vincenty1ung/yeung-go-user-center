package logic

import (
	"context"

	"github.com/uncleyeung/yeung-go-user-center/rpc/user/internal/svc"
	"github.com/uncleyeung/yeung-go-user-center/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserAccountSelectByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAccountSelectByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAccountSelectByUserIdLogic {
	return &UserAccountSelectByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAccountSelectByUserIdLogic) UserAccountSelectByUserId(in *user.GetByIdReq) (*user.AccountResp, error) {
	// todo: add your logic here and delete this line

	return &user.AccountResp{}, nil
}
