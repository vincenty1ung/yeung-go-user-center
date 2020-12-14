package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/uncleyeung/yeung-user-center/rpc/user/internal/config"
	"github.com/uncleyeung/yeung-user-center/rpc/user/internal/db"
)

type ServiceContext struct {
	c             config.Config
	UserDb        db.UserModel
	UserAccountDb db.UserAccountModel
	UserAddressDb db.UserAddressModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c:             c,
		UserDb:        db.NewUserModel(sqlx.NewMysql(c.DataSource), c.Cache),
		UserAccountDb: db.NewUserAccountModel(sqlx.NewMysql(c.DataSource), c.Cache),
		UserAddressDb: db.NewUserAddressModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
