package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	userAddressFieldNames          = builderx.FieldNames(&UserAddress{})
	userAddressRows                = strings.Join(userAddressFieldNames, ",")
	userAddressRowsExpectAutoSet   = strings.Join(stringx.Remove(userAddressFieldNames, "id", "create_time", "update_time"), ",")
	userAddressRowsWithPlaceHolder = strings.Join(stringx.Remove(userAddressFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"

	cacheUserAddressIdPrefix = "cache#UserAddress#id#"
)

type (
	UserAddressModel interface {
		Insert(data UserAddress) (sql.Result, error)
		FindOne(id int64) (*UserAddress, error)
		Update(data UserAddress) error
		Delete(id int64) error
	}

	defaultUserAddressModel struct {
		sqlc.CachedConn
		table string
	}

	UserAddress struct {
		Id      int64  `db:"id"`      // 主键id
		UserId  int64  `db:"user_id"` // 用户id
		Address string `db:"address"` // 地址
		Status  int64  `db:"status"`  // 0:默认1:非默认
	}
)

func NewUserAddressModel(conn sqlx.SqlConn, c cache.CacheConf) UserAddressModel {
	return &defaultUserAddressModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "user_address",
	}
}

func (m *defaultUserAddressModel) Insert(data UserAddress) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userAddressRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.UserId, data.Address, data.Status)

	return ret, err
}

func (m *defaultUserAddressModel) FindOne(id int64) (*UserAddress, error) {
	userAddressIdKey := fmt.Sprintf("%s%v", cacheUserAddressIdPrefix, id)
	var resp UserAddress
	err := m.QueryRow(&resp, userAddressIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where id = ? limit 1", userAddressRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserAddressModel) Update(data UserAddress) error {
	userAddressIdKey := fmt.Sprintf("%s%v", cacheUserAddressIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = ?", m.table, userAddressRowsWithPlaceHolder)
		return conn.Exec(query, data.UserId, data.Address, data.Status, data.Id)
	}, userAddressIdKey)
	return err
}

func (m *defaultUserAddressModel) Delete(id int64) error {

	userAddressIdKey := fmt.Sprintf("%s%v", cacheUserAddressIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = ?", m.table)
		return conn.Exec(query, id)
	}, userAddressIdKey)
	return err
}

func (m *defaultUserAddressModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserAddressIdPrefix, primary)
}

func (m *defaultUserAddressModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", userAddressRows, m.table)
	return conn.QueryRow(v, query, primary)
}
