// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CityVisitor is the golang structure of table city_visitor for DAO operations like Where/Data.
type CityVisitor struct {
	g.Meta `orm:"table:city_visitor, do:true"`
	City   interface{} // 城市名称
	Uv     interface{} // 独立访客数量
}
