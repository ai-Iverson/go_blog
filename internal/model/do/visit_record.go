// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// VisitRecord is the golang structure of table visit_record for DAO operations like Where/Data.
type VisitRecord struct {
	g.Meta `orm:"table:visit_record, do:true"`
	Id     interface{} //
	Pv     interface{} // 访问量
	Uv     interface{} // 独立用户
	Date   interface{} // 日期"02-23"
}
