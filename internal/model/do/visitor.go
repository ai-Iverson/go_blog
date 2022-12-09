// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Visitor is the golang structure of table visitor for DAO operations like Where/Data.
type Visitor struct {
	g.Meta     `orm:"table:visitor, do:true"`
	Id         interface{} //
	Uuid       interface{} // 访客标识码
	Ip         interface{} // ip
	IpSource   interface{} // ip来源
	Os         interface{} // 操作系统
	Browser    interface{} // 浏览器
	CreateTime *gtime.Time // 首次访问时间
	LastTime   *gtime.Time // 最后访问时间
	Pv         interface{} // 访问页数统计
	UserAgent  interface{} // user-agent用户代理
}
