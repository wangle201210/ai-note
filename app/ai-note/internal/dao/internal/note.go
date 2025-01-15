// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NoteDao is the data access object for the table note.
type NoteDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of the current DAO.
	columns NoteColumns // columns contains all the column names of Table for convenient usage.
}

// NoteColumns defines and stores column names for the table note.
type NoteColumns struct {
	Id         string // note id
	UserId     string // user id
	CategoryId string // category id
	Amount     string // 具体金额
	OrgContent string // 原始内容
	Status     string // note status
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// noteColumns holds the columns for the table note.
var noteColumns = NoteColumns{
	Id:         "id",
	UserId:     "user_id",
	CategoryId: "category_id",
	Amount:     "amount",
	OrgContent: "org_content",
	Status:     "status",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewNoteDao creates and returns a new DAO object for table data access.
func NewNoteDao() *NoteDao {
	return &NoteDao{
		group:   "default",
		table:   "note",
		columns: noteColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *NoteDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *NoteDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *NoteDao) Columns() NoteColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *NoteDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *NoteDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *NoteDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
