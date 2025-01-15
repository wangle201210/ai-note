// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Note is the golang structure for table note.
type Note struct {
	Id         uint        `json:"id"         orm:"id"          description:"note id"`     // note id
	UserId     int         `json:"userId"     orm:"user_id"     description:"user id"`     // user id
	CategoryId int         `json:"categoryId" orm:"category_id" description:"category id"` // category id
	Amount     float64     `json:"amount"     orm:"amount"      description:"具体金额"`        // 具体金额
	OrgContent string      `json:"orgContent" orm:"org_content" description:"原始内容"`        // 原始内容
	Status     int         `json:"status"     orm:"status"      description:"note status"` // note status
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`        // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"更新时间"`        // 更新时间
}
