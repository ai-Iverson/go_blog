package v1

import "github.com/gogf/gf/v2/frame/g"

type SiteSettingsDetilReq struct {
	g.Meta        `path:"/admin/siteSettings" method:"get" summart:" 获取站点信息" tags:"获取站点信息"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
}

type SiteSettingsDetilRes struct {
	Type3 []Type3 `json:"type3"`
	Type2 []Type2 `json:"type2"`
	Type1 []Type1 `json:"type1"`
}
type Type3 struct {
	ID     int    `json:"id"`
	NameEn string `json:"nameEn"`
	NameZh string `json:"nameZh"`
	Value  string `json:"value"`
	Type   int    `json:"type"`
}
type Type2 struct {
	ID     int    `json:"id"`
	NameEn string `json:"nameEn"`
	NameZh string `json:"nameZh"`
	Value  string `json:"value"`
	Type   int    `json:"type"`
}
type Type1 struct {
	ID     int    `json:"id"`
	NameEn string `json:"nameEn"`
	NameZh string `json:"nameZh"`
	Value  string `json:"value"`
	Type   int    `json:"type"`
}

type UpdateSiteSettingsReq struct {
	g.Meta        `path:"/admin/siteSettings" method:"post" summart:" 获取站点信息" tags:"获取站点信息"`
	Authorization string     `json:"Authorization" in:"header"  dc:"Authorization"`
	Settings      []Settings `json:"settings"`
	DeleteIds     []int      `json:"deleteIds"`
}

type Settings struct {
	ID     int    `json:"id"`
	NameEn string `json:"nameEn"`
	NameZh string `json:"nameZh"`
	Value  string `json:"value"`
	Type   int    `json:"type"`
}

type UpdateSiteSettingsRes struct {
}
