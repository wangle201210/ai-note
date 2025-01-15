package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HelloReq struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"You first hello api"`
	Amount int64  `json:"amount"`
	Type   string `json:"type"`
	Org    string `json:"org"`
}

type HelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
