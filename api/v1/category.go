package v1

import "github.com/gogf/gf/v2/frame/g"

type ShowCategoryReq struct {
	g.Meta   `path:"/admin/categories" method:"get" summart:"获取标签页" tags:"获取标签页"`
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

type ShowCategoryRes struct {
	List []Categorys `json:"list"`
	CommonPaginationReq
}

type Categorys struct {
	Id           int      `json:"id"`
	CategoryName string   `json:"name"`
	Blogs        []string `json:"blogs"`
}

type AddCategoryReq struct {
	g.Meta       `path:"/admin/category" method:"post" summart:"添加标签" tags:"添加标签"`
	CategoryName string `json:"name"`
	Color        string `json:"color"`
}

type AddCategoryRes struct {
}

type UpdateCategoryReq struct {
	g.Meta       `path:"/admin/category" method:"put" summart:"更新标签" tags:"更新标签"`
	Id           int    `json:"id"`
	CategoryName string `json:"name"`
	Color        string `json:"color"`
}

type UpdateCategoryRes struct {
}

type DeleteCategoryReq struct {
	g.Meta `path:"/admin/category" method:"delete" summart:"删除标签" tags:"删除标签"`
	Id     int `json:"id"`
}

type DeleteCategoryRes struct {
}
