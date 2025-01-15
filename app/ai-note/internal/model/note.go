package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model/entity"
)

const (
	AllCategory = "all"
)

type NoteInfo struct {
	*entity.Note
	User     *entity.User     `orm:"with:id=user_id" json:"user"`
	Category *entity.Category `orm:"with:id=category_id" json:"category"`
}

type NoteSimpleInfo struct {
	Id           uint        `json:"id,omitempty"         orm:"id"          description:"note id" ` // note id
	Amount       float64     `json:"amount"     orm:"amount"      description:"具体金额"`               // 具体金额
	CreatedAt    *gtime.Time `json:"createdAt,omitempty"  orm:"created_at"  description:"创建时间"`     // 创建时间
	UserName     string      `json:"user_name,omitempty"`
	CategoryName string      `json:"category_name"`
}
