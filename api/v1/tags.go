package v1

import "github.com/gogf/gf/v2/frame/g"

type ShowTagsReq struct {
	g.Meta   `path:"/admin/tags" method:"get" summart:"获取标签页" tags:"获取标签页"`
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

type ShowTagsRes struct {
	List []Tags `json:"list"`
	CommonPaginationReq
}

type Tags struct {
	Id      int      `json:"id"`
	TagName string   `json:"name"`
	Color   string   `json:"color"`
	Blogs   []string `json:"blogs"`
}

type AddTagsReq struct {
	g.Meta  `path:"/admin/tag" method:"post" summart:"添加标签" tags:"添加标签"`
	TagName string `json:"name"`
	Color   string `json:"color"`
}

type AddTagsRes struct {
}

type UpdateTagsReq struct {
	g.Meta  `path:"/admin/tag" method:"put" summart:"更新标签" tags:"更新标签"`
	Id      int    `json:"id"`
	TagName string `json:"name"`
	Color   string `json:"color"`
}

type UpdateTagsRes struct {
}

type DeleteTagsReq struct {
	g.Meta `path:"/admin/tag" method:"delete" summart:"删除标签" tags:"删除标签"`
	Id     int `json:"id"`
}

type DeleteTagsRes struct {
}
