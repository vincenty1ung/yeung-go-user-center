package logic

import (
	"context"
	"github.com/ulule/deepcopier"

	"github.com/uncleyeung/yeung-go-user-center/rpc/user/internal/svc"
	"github.com/uncleyeung/yeung-go-user-center/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserSelectByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserSelectByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSelectByIdLogic {
	return &UserSelectByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserSelectByIdLogic) UserSelectById(in *user.GetByIdReq) (*user.UserGetByIdResp, error) {
	dbs := l.svcCtx.UserDb
	one, err := dbs.FindOne(in.Id)
	if err != nil {
		l.Error(err)
		return nil, err
	}
	_, err = l.svcCtx.Adder.Add(l.ctx, nil)
	l.Logger.Error(err)
	var result user.UserGetByIdResp
	_ = deepcopier.Copy(one).To(&result)
	result.Status = CheckState(one.Status)
	return &result, nil
}
func CheckState(status int32) user.UserGetByIdResp_Status {
	if status == 0 {
		return user.UserGetByIdResp_INIT
	}
	if status == 1 {
		return user.UserGetByIdResp_ACTIVATION
	}
	if status == 2 {
		return user.UserGetByIdResp_LOGOUT
	}
	return -1
}
