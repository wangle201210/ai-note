package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model/entity"
)

// Status marks user status.
type Status int

const (
	StatusOK       Status = 0 // User is OK.
	StatusDisabled Status = 1 // User is disabled.
)

type CreateReq struct {
	g.Meta `path:"/user" method:"post" tags:"User" summary:"Create user"`
	Name   string `v:"required" dc:"user name"`
	OpenId string `v:"required" dc:"user openId"`
}

type CreateRes struct {
	Id int64 `json:"id" dc:"user id"`
}

type UpdateReq struct {
	g.Meta `path:"/user/{id}" method:"put" tags:"User" summary:"Update user"`
	Id     int64   `v:"required" dc:"user id"`
	Name   *string `v:"length:3,10" dc:"user name"`
	OpenId *string `v:"required" dc:"user openID"`
	Status *Status `v:"in:0,1" dc:"user status"`
}
type UpdateRes struct{}

type DeleteReq struct {
	g.Meta `path:"/user/{id}" method:"delete" tags:"User" summary:"Delete user"`
	Id     int64 `v:"required" dc:"user id"`
}
type DeleteRes struct{}

type GetOneReq struct {
	g.Meta `path:"/user/{id}" method:"get" tags:"User" summary:"Get one user"`
	Id     int64 `v:"required" dc:"user id"`
}
type GetOneRes struct {
	*entity.User `dc:"user"`
}

type GetListReq struct {
	g.Meta `path:"/user" method:"get" tags:"User" summary:"Get users"`
	Name   *string `dc:"user name"`
	OpenId *string `dc:"user openID"`
	Status *Status `v:"in:0,1" dc:"user age"`
}
type GetListRes struct {
	List []*entity.User `json:"list" dc:"user list"`
}

type GetOrSaveReq struct {
	g.Meta `path:"/user/get_or_save" method:"get" tags:"User" summary:"Create user"`
	Name   string `v:"required" dc:"user name"`
	OpenId string `dc:"user openId"`
}

type GetOrSaveRes struct {
	Name   string `dc:"user name"`
	OpenId string `dc:"user openId"`
}
