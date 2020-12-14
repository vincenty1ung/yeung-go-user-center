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
	userAccountFieldNames          = builderx.FieldNames(&UserAccount{})
	userAccountRows                = strings.Join(userAccountFieldNames, ",")
	userAccountRowsExpectAutoSet   = strings.Join(stringx.Remove(userAccountFieldNames, "id", "create_time", "update_time"), ",")
	userAccountRowsWithPlaceHolder = strings.Join(stringx.Remove(userAccountFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"

	cacheUserAccountIdPrefix = "cache#UserAccount#id#"
)

type (
	UserAccountModel interface {
		Insert(data UserAccount) (sql.Result, error)
		FindOne(id int64) (*UserAccount, error)
		Update(data UserAccount) error
		Delete(id int64) error
	}

	defaultUserAccountModel struct {
		sqlc.CachedConn
		table string
	}

	UserAccount struct {
		Id      int64  `db:"id"`      // 主键id
		UserId  int64  `db:"user_id"` // 用户id
		Balance string `db:"balance"` // 余额
		Loan    string `db:"loan"`    // 贷款
	}
)

func NewUserAccountModel(conn sqlx.SqlConn, c cache.CacheConf) UserAccountModel {
	return &defaultUserAccountModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "user_account",
	}
}

func (m *defaultUserAccountModel) Insert(data UserAccount) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userAccountRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.UserId, data.Balance, data.Loan)

	return ret, err
}

func (m *defaultUserAccountModel) FindOne(id int64) (*UserAccount, error) {
	userAccountIdKey := fmt.Sprintf("%s%v", cacheUserAccountIdPrefix, id)
	var resp UserAccount
	err := m.QueryRow(&resp, userAccountIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where id = ? limit 1", userAccountRows, m.table)
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

func (m *defaultUserAccountModel) Update(data UserAccount) error {
	userAccountIdKey := fmt.Sprintf("%s%v", cacheUserAccountIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = ?", m.table, userAccountRowsWithPlaceHolder)
		return conn.Exec(query, data.UserId, data.Balance, data.Loan, data.Id)
	}, userAccountIdKey)
	return err
}

func (m *defaultUserAccountModel) Delete(id int64) error {

	userAccountIdKey := fmt.Sprintf("%s%v", cacheUserAccountIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = ?", m.table)
		return conn.Exec(query, id)
	}, userAccountIdKey)
	return err
}

func (m *defaultUserAccountModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserAccountIdPrefix, primary)
}

func (m *defaultUserAccountModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", userAccountRows, m.table)
	return conn.QueryRow(v, query, primary)
}
