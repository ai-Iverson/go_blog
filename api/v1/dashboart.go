package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DashboartReq struct {
	g.Meta        `path:"/admin/dashboard" method:"get" summart:"界面" tags:"界面"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
}

type DashboardRes struct {
	Uv           int           `json:"uv"`
	BlogCount    int           `json:"blogCount"`
	Pv           int           `json:"pv"`
	CityVisitor  []CityVisitor `json:"cityVisitor"`
	Tag          Tag           `json:"tag"`
	Category     Category      `json:"category"`
	VisitRecord  VisitRecord   `json:"visitRecord"`
	CommentCount int           `json:"commentCount"`
}

type CityVisitor struct {
	City string `json:"city"`
	Uv   int    `json:"uv"`
}
type Series struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}
type Tag struct {
	Legend []string `json:"legend"`
	Series []Series `json:"series"`
}
type Category struct {
	Legend []string `json:"legend"`
	Series []Series `json:"series"`
}
type VisitRecord struct {
	Date []string `json:"date"`
	Uv   []int    `json:"uv"`
	Pv   []int    `json:"pv"`
}
