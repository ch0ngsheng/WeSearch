package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ QueryLogsModel = (*customQueryLogsModel)(nil)

type (
	// QueryLogsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customQueryLogsModel.
	QueryLogsModel interface {
		queryLogsModel
	}

	customQueryLogsModel struct {
		*defaultQueryLogsModel
	}
)

// NewQueryLogsModel returns a model for the database table.
func NewQueryLogsModel(conn sqlx.SqlConn) QueryLogsModel {
	return &customQueryLogsModel{
		defaultQueryLogsModel: newQueryLogsModel(conn),
	}
}
