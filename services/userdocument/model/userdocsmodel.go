package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserDocsModel = (*customUserDocsModel)(nil)

type (
	// UserDocsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserDocsModel.
	UserDocsModel interface {
		userDocsModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
	}

	customUserDocsModel struct {
		*defaultUserDocsModel
	}
)

// NewUserDocsModel returns a model for the database table.
func NewUserDocsModel(conn sqlx.SqlConn) UserDocsModel {
	return &customUserDocsModel{
		defaultUserDocsModel: newUserDocsModel(conn),
	}
}

// Trans wraps fn in database transaction.
func (m *defaultUserDocsModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}
