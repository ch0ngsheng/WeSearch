package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DocumentsModel = (*customDocumentsModel)(nil)

type (
	// DocumentsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDocumentsModel.
	DocumentsModel interface {
		documentsModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		FindByUrlHash(ctx context.Context, session sqlx.Session, urlHash string) ([]*Documents, error)
		FindByUID(ctx context.Context, session sqlx.Session, userID int64) ([]*Documents, error)
		FindOneByUIDAndDocID(ctx context.Context, session sqlx.Session, userID int64, docID int64) (*Documents, error)
	}

	customDocumentsModel struct {
		*defaultDocumentsModel
	}
)

// NewDocumentsModel returns a model for the database table.
func NewDocumentsModel(conn sqlx.SqlConn) DocumentsModel {
	return &customDocumentsModel{
		defaultDocumentsModel: newDocumentsModel(conn),
	}
}

// Trans wraps fn in database transaction.
func (m *defaultDocumentsModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

func (m *customDocumentsModel) FindByUrlHash(ctx context.Context, session sqlx.Session, urlHash string) ([]*Documents, error) {
	var resp []*Documents
	query := fmt.Sprintf("select %s from %s where `hash` = ?", documentsRows, m.table)
	var err error
	if session != nil {
		err = session.QueryRowsCtx(ctx, &resp, query, urlHash)
	} else {
		err = m.conn.QueryRowsCtx(ctx, &resp, query, urlHash)
	}
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customDocumentsModel) FindByUID(ctx context.Context, session sqlx.Session, userID int64) ([]*Documents, error) {
	var resp []*Documents

	query := fmt.Sprintf(
		"select d.id, d.url, d.hash, d.title, d.description, ud.created_at, ud.created_at from %s as d "+
			"left join "+
			"%s as ud on "+
			"ud.doc_id=d.id where ud.uid= ?",
		m.tableName(), newUserDocsModel(nil).tableName(),
	)

	var err error
	if session != nil {
		err = session.QueryRowsCtx(ctx, &resp, query, userID)
	} else {
		err = m.conn.QueryRowsCtx(ctx, &resp, query, userID)
	}

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customDocumentsModel) FindOneByUIDAndDocID(ctx context.Context, session sqlx.Session, userID int64, docID int64) (*Documents, error) {
	var resp Documents

	query := fmt.Sprintf(
		"select d.id, d.url, d.hash, d.title, d.description, ud.created_at, ud.created_at from %s as d "+
			"left join "+
			"%s as ud on "+
			"ud.doc_id=d.id where ud.uid= ? and d.id= ?",
		m.tableName(), newUserDocsModel(nil).tableName(),
	)

	var err error
	if session != nil {
		err = session.QueryRowCtx(ctx, &resp, query, userID, docID)
	} else {
		err = m.conn.QueryRowCtx(ctx, &resp, query, userID, docID)
	}
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, err
	default:
		return nil, err
	}
}
