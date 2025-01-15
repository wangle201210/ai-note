// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Category is the golang structure of table category for DAO operations like Where/Data.
type Category struct {
	g.Meta `orm:"table:category, do:true"`
	Id     interface{} // category id
	Name   interface{} // category name
	Status interface{} // category status
}
