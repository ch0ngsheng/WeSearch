package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		UpdateTimeByID(ctx context.Context, session sqlx.Session, newData *Users) (sql.Result, error)
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn),
	}
}

// Trans wraps fn in database transaction.
func (m *defaultUsersModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

func (m *customUsersModel) UpdateTimeByID(ctx context.Context, session sqlx.Session, newData *Users) (sql.Result, error) {
	query := fmt.Sprintf("update %s set updated_time=? where `id` = ?", m.table)
	if session != nil {
		return session.ExecCtx(ctx, query, newData.UpdatedAt, newData.Id)
	}
	return m.conn.ExecCtx(ctx, query, newData.UpdatedAt, newData.Id)
}
