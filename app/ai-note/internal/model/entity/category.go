// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Category is the golang structure for table category.
type Category struct {
	Id     uint   `json:"id"     orm:"id"     description:"category id"`     // category id
	Name   string `json:"name"   orm:"name"   description:"category name"`   // category name
	Status int    `json:"status" orm:"status" description:"category status"` // category status
}
