package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model/entity"
)

// Status marks category status.
type Status int

const (
	StatusOK       Status = 0 // User is OK.
	StatusDisabled Status = 1 // User is disabled.
)

type CreateReq struct {
	g.Meta `path:"/category" method:"post" tags:"User" summary:"Create category"`
	Name   string `v:"required|length:3,10" dc:"category name"`
}
type CreateRes struct {
	Id int64 `json:"id" dc:"category id"`
}

type UpdateReq struct {
	g.Meta `path:"/category/{id}" method:"put" tags:"User" summary:"Update category"`
	Id     int64   `v:"required" dc:"category id"`
	Name   *string `v:"length:3,10" dc:"category name"`
	Status *Status `v:"in:0,1" dc:"category status"`
}
type UpdateRes struct{}

type DeleteReq struct {
	g.Meta `path:"/category/{id}" method:"delete" tags:"User" summary:"Delete category"`
	Id     int64 `v:"required" dc:"category id"`
}
type DeleteRes struct{}

type GetOneReq struct {
	g.Meta `path:"/category/{id}" method:"get" tags:"User" summary:"Get one category"`
	Id     int64 `v:"required" dc:"category id"`
}
type GetOneRes struct {
	*entity.Category `dc:"category"`
}

type GetListReq struct {
	g.Meta `path:"/category" method:"get" tags:"User" summary:"Get categorys"`
	Name   *string `dc:"category name"`
	Status *Status `v:"in:0,1" dc:"category age"`
}

type GetListRes struct {
	List []*entity.Category `json:"list" dc:"category list"`
}

type GetCategoryNameListReq struct {
	g.Meta `path:"/category/list" method:"get" tags:"User" summary:"Get categorys"`
}

type GetCategoryNameListRes struct {
	Name []string `json:"name" dc:"category list"`
}
