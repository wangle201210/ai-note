package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model/entity"
	"time"
)

// Status marks user status.
type Status int

const (
	StatusOK       Status = 0 // User is OK.
	StatusDisabled Status = 1 // User is disabled.
)

type CreateReq struct {
	g.Meta     `path:"/note" method:"post" tags:"Note" summary:"Create note"`
	UserId     int64   `v:"required" dc:"user id"`
	CategoryId int64   `v:"required" dc:"category id"`
	Amount     float64 `v:"required" dc:"amount"`
	OrgContent string  `v:"required"`
}
type CreateRes struct {
	Id int64 `json:"id" dc:"user id"`
}

type UpdateReq struct {
	g.Meta     `path:"/note/{id}" method:"put" tags:"Note" summary:"Update note"`
	Id         int64 `v:"required" dc:"user id"`
	UserId     *int64
	CategoryId *int64
	Amount     *float64
	OrgContent *string
	Status     *Status `v:"in:0,1" dc:"user status"`
}
type UpdateRes struct{}

type DeleteReq struct {
	g.Meta `path:"/note/{id}" method:"delete" tags:"Note" summary:"Delete note"`
	Id     int64 `v:"required" dc:"user id"`
}
type DeleteRes struct{}

type GetOneReq struct {
	g.Meta `path:"/note/{id}" method:"get" tags:"Note" summary:"Get one note"`
	Id     int64 `v:"required" dc:"user id"`
}
type GetOneRes struct {
	*entity.Note `dc:"note"`
}

type GetListReq struct {
	g.Meta     `path:"/note" method:"get" tags:"Note" summary:"Get note"`
	UserId     *int64
	CategoryId *int64
	Status     *Status `v:"in:0,1" dc:"user age"`
}
type GetListRes struct {
	List []*entity.Note `json:"list" dc:"user list"`
}

type AddNoteReq struct {
	g.Meta       `path:"/note/add" method:"post" tags:"Note" summary:"Create note"`
	UserName     string  `json:"user_name" v:"required" dc:"user name"`
	CategoryName string  `json:"category_name" v:"required" dc:"category name"`
	Amount       float64 `json:"amount" v:"required" dc:"amount"`
	OrgContent   string  `json:"org_content"v:"required"`
}

type AddNoteRes struct {
	Id int64 `json:"id" dc:"note id"`
}

type GetNoteListReq struct {
	g.Meta   `path:"/note/list" method:"get" tags:"Note" summary:"Get note"`
	UserName string `json:"user_name"`
}

type GetNoteListRes struct {
	List []*model.NoteSimpleInfo `json:"list" dc:"user list"`
}

type GetNoteInfoReq struct {
	g.Meta       `path:"/note/info" method:"get" tags:"Note" summary:"Get note"`
	UserName     string     `json:"user_name"`
	Start        *time.Time `json:"start"`
	End          *time.Time `json:"end"`
	CategoryName string     `json:"category"`
}

type GetNoteInfoRes struct {
	Amount float64                 `json:"amount" dc:"amount"`
	List   []*model.NoteSimpleInfo `json:"list" dc:"user list"`
}
