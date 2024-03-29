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
	queryLogsFieldNames          = builder.RawFieldNames(&QueryLogs{})
	queryLogsRows                = strings.Join(queryLogsFieldNames, ",")
	queryLogsRowsExpectAutoSet   = strings.Join(stringx.Remove(queryLogsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	queryLogsRowsWithPlaceHolder = strings.Join(stringx.Remove(queryLogsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	queryLogsModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *QueryLogs) (sql.Result, error)
		FindOne(ctx context.Context, session sqlx.Session, id int64) (*QueryLogs, error)
		Update(ctx context.Context, session sqlx.Session, data *QueryLogs) (sql.Result, error)
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultQueryLogsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	QueryLogs struct {
		Id        int64     `db:"id"`         // log id
		Uid       int64     `db:"uid"`        // user id
		Content   string    `db:"content"`    // query content
		CreatedAt time.Time `db:"created_at"` // query time
	}
)

func newQueryLogsModel(conn sqlx.SqlConn) *defaultQueryLogsModel {
	return &defaultQueryLogsModel{
		conn:  conn,
		table: "`query_logs`",
	}
}

func (m *defaultQueryLogsModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	if session != nil {
		_, err := session.ExecCtx(ctx, query, id)
		return err
	}
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultQueryLogsModel) FindOne(ctx context.Context, session sqlx.Session, id int64) (*QueryLogs, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", queryLogsRows, m.table)
	var resp QueryLogs
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

func (m *defaultQueryLogsModel) Insert(ctx context.Context, session sqlx.Session, data *QueryLogs) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, queryLogsRowsExpectAutoSet)
	if session != nil {
		return session.ExecCtx(ctx, query, data.Uid, data.Content)
	}
	return m.conn.ExecCtx(ctx, query, data.Uid, data.Content)
}

func (m *defaultQueryLogsModel) Update(ctx context.Context, session sqlx.Session, data *QueryLogs) (sql.Result, error) {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, queryLogsRowsWithPlaceHolder)
	if session != nil {
		return session.ExecCtx(ctx, query, data.Uid, data.Content, data.Id)
	}
	return m.conn.ExecCtx(ctx, query, data.Uid, data.Content, data.Id)
}

func (m *defaultQueryLogsModel) tableName() string {
	return m.table
}
