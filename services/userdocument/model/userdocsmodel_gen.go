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
	userDocsFieldNames          = builder.RawFieldNames(&UserDocs{})
	userDocsRows                = strings.Join(userDocsFieldNames, ",")
	userDocsRowsExpectAutoSet   = strings.Join(stringx.Remove(userDocsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userDocsRowsWithPlaceHolder = strings.Join(stringx.Remove(userDocsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userDocsModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *UserDocs) (sql.Result, error)
		FindOne(ctx context.Context, session sqlx.Session, id int64) (*UserDocs, error)
		Update(ctx context.Context, session sqlx.Session, data *UserDocs) (sql.Result, error)
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultUserDocsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserDocs struct {
		Id        int64     `db:"id"`         // id
		Uid       int64     `db:"uid"`        // user id
		DocId     int64     `db:"doc_id"`     // document id
		CreatedAt time.Time `db:"created_at"` // user collect time
	}
)

func newUserDocsModel(conn sqlx.SqlConn) *defaultUserDocsModel {
	return &defaultUserDocsModel{
		conn:  conn,
		table: "`user_docs`",
	}
}

func (m *defaultUserDocsModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	if session != nil {
		_, err := session.ExecCtx(ctx, query, id)
		return err
	}
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserDocsModel) FindOne(ctx context.Context, session sqlx.Session, id int64) (*UserDocs, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userDocsRows, m.table)
	var resp UserDocs
	var err error
	if session != nil {
		err = session.QueryRowCtx(ctx, &resp, query, id)
	} else {
		err = m.conn.QueryRowCtx(ctx, &resp, query, id)
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

func (m *defaultUserDocsModel) Insert(ctx context.Context, session sqlx.Session, data *UserDocs) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, userDocsRowsExpectAutoSet)
	if session != nil {
		return session.ExecCtx(ctx, query, data.Uid, data.DocId)
	}
	return m.conn.ExecCtx(ctx, query, data.Uid, data.DocId)
}

func (m *defaultUserDocsModel) Update(ctx context.Context, session sqlx.Session, data *UserDocs) (sql.Result, error) {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userDocsRowsWithPlaceHolder)
	if session != nil {
		return session.ExecCtx(ctx, query, data.Uid, data.DocId, data.Id)
	}
	return m.conn.ExecCtx(ctx, query, data.Uid, data.DocId, data.Id)
}

func (m *defaultUserDocsModel) tableName() string {
	return m.table
}
