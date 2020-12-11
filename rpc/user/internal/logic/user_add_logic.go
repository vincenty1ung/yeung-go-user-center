package logic

import (
	"context"
	"github.com/ulule/deepcopier"
	db2 "github.com/uncleyeung/yeung-user-center/rpc/user/db"
	"time"

	"github.com/uncleyeung/yeung-user-center/rpc/user/internal/svc"
	"github.com/uncleyeung/yeung-user-center/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAddLogic) UserAdd(in *user.UserAddReq) (*user.UserAddResp, error) {
	// todo: add your logic here and delete this line
	db := l.svcCtx.UserDb
	var userDbTmp db2.User
	resp := user.UserAddResp{}
	_ = deepcopier.Copy(in).To(&userDbTmp)
	userDbTmp.UpdateTime = time.Now().Unix()
	userDbTmp.CreateTime = time.Now().Unix()
	insert, err := db.Insert(userDbTmp)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	id, _ := insert.LastInsertId()
	l.Logger.Infof("数据结果id:%d", id)
	resp.Ok = true
	return &resp, nil
}
