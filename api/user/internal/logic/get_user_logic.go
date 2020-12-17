package logic

import (
	"context"
	"github.com/ulule/deepcopier"
	"github.com/uncleyeung/yeung-go-user-center/rpc/user/user"

	"github.com/uncleyeung/yeung-go-user-center/api/user/internal/svc"
	"github.com/uncleyeung/yeung-go-user-center/api/user/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserLogic {
	return GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req types.GetByIdReq) (*types.UserGetByIdResp, error) {
	service := l.svcCtx.UserService
	idReq := user.GetByIdReq{Id: req.Id}
	id, err := service.UserSelectById(l.ctx, &idReq)
	if err != nil {
		l.Error(err)
		return nil, err
	}
	var result types.UserGetByIdResp
	_ = deepcopier.Copy(id).To(&result)
	result.Status = CheckState(id.Status)
	return &result, nil
}
func CheckState(status user.UserGetByIdResp_Status) types.UserStatus {
	if status == 0 {
		return types.UserAddReq_INIT
	}
	if status == 1 {
		return types.UserAddReq_ACTIVATION
	}
	if status == 2 {
		return types.UserAddReq_LOGOUT
	}
	return -1
}
