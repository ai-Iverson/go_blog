// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VisitRecordDao is the data access object for table visit_record.
type VisitRecordDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns VisitRecordColumns // columns contains all the column names of Table for convenient usage.
}

// VisitRecordColumns defines and stores column names for table visit_record.
type VisitRecordColumns struct {
	Id   string //
	Pv   string // 访问量
	Uv   string // 独立用户
	Date string // 日期"02-23"
}

//  visitRecordColumns holds the columns for table visit_record.
var visitRecordColumns = VisitRecordColumns{
	Id:   "id",
	Pv:   "pv",
	Uv:   "uv",
	Date: "date",
}

// NewVisitRecordDao creates and returns a new DAO object for table data access.
func NewVisitRecordDao() *VisitRecordDao {
	return &VisitRecordDao{
		group:   "default",
		table:   "visit_record",
		columns: visitRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *VisitRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *VisitRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *VisitRecordDao) Columns() VisitRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *VisitRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *VisitRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *VisitRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
