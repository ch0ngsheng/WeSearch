// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	usersFieldNames          = builder.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	usersModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *Users) (sql.Result, error)
		FindOne(ctx context.Context, session sqlx.Session, id int64) (*Users, error)
		FindOneByOpenid(ctx context.Context, session sqlx.Session, openid string) (*Users, error)
		Update(ctx context.Context, session sqlx.Session, data *Users) (sql.Result, error)
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultUsersModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Users struct {
		Id        int64     `db:"id"`         // user id
		Openid    string    `db:"openid"`     // wechat openid
		CreatedAt time.Time `db:"created_at"` // user first collect time
		UpdatedAt time.Time `db:"updated_at"` // user last collect time
	}
)

func newUsersModel(conn sqlx.SqlConn) *defaultUsersModel {
	return &defaultUsersModel{
		conn:  conn,
		table: "`users`",
	}
}

func (m *defaultUsersModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	if session != nil {
		_, err := session.ExecCtx(ctx, query, id)
		return err
	}
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUsersModel) FindOne(ctx context.Context, session sqlx.Session, id int64) (*Users, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
	var resp Users
	var err error
	if session != nil {
		err = session.QueryRowCtx(ctx, &resp, query, id)
	}
	err = m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByOpenid(ctx context.Context, session sqlx.Session, openid string) (*Users, error) {
	var resp Users
	query := fmt.Sprintf("select %s from %s where `openid` = ? limit 1", usersRows, m.table)
	var err error
	if session != nil {
		err = session.QueryRowCtx(ctx, &resp, query, openid)
	} else {
		err = m.conn.QueryRowCtx(ctx, &resp, query, openid)
	}

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Insert(ctx context.Context, session sqlx.Session, data *Users) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?)", m.table, usersRowsExpectAutoSet)
	if session != nil {
		return session.ExecCtx(ctx, query, data.Openid)
	}
	return m.conn.ExecCtx(ctx, query, data.Openid)
}

func (m *defaultUsersModel) Update(ctx context.Context, session sqlx.Session, newData *Users) (sql.Result, error) {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersRowsWithPlaceHolder)
	if session != nil {
		return session.ExecCtx(ctx, query, newData.Openid, newData.Id)
	}
	return m.conn.ExecCtx(ctx, query, newData.Openid, newData.Id)
}

func (m *defaultUsersModel) tableName() string {
	return m.table
}
