// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// User is the golang structure for table user.
type User struct {
	Id     uint   `json:"id"     orm:"id"      description:"user id"`        // user id
	Name   string `json:"name"   orm:"name"    description:"user name"`      // user name
	Status int    `json:"status" orm:"status"  description:"user status"`    // user status
	OpenId string `json:"openId" orm:"open_id" description:"wechat open_id"` // wechat open_id
}
