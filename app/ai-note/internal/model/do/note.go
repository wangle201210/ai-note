// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Note is the golang structure of table note for DAO operations like Where/Data.
type Note struct {
	g.Meta     `orm:"table:note, do:true"`
	Id         interface{} // note id
	UserId     interface{} // user id
	CategoryId interface{} // category id
	Amount     interface{} // 具体金额
	OrgContent interface{} // 原始内容
	Status     interface{} // note status
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
