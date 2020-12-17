package logic

import (
	"context"
	"github.com/ulule/deepcopier"
	"github.com/uncleyeung/yeung-go-user-center/rpc/user/user"

	"github.com/uncleyeung/yeung-go-user-center/api/user/internal/svc"
	"github.com/uncleyeung/yeung-go-user-center/api/user/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateUserLogic {
	return CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req types.UserAddRequest) (*types.UserGetByIdResp, error) {
	userService := l.svcCtx.UserService
	var user user.UserAddReq
	var resp types.UserGetByIdResp
	_ = deepcopier.Copy(req).To(&user)
	l.Infof("数据%+v", user)
	add, err := userService.UserAdd(l.ctx, &user)
	if err != nil {
		return nil, err
	}
	if add.Ok {
		_ = deepcopier.Copy(req).To(&resp)
	}
	return &resp, nil
}
