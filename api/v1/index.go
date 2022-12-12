package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CommonPaginationReq struct {
	Total             int   `json:"total"`
	PageNum           int   `json:"pageNum"`
	PageSize          int   `json:"pageSize"`
	Size              int   `json:"size"`
	StartRow          int   `json:"startRow"`
	EndRow            int   `json:"endRow"`
	Pages             int   `json:"pages"`
	PrePage           int   `json:"prePage"`
	NextPage          int   `json:"nextPage"`
	IsFirstPage       bool  `json:"isFirstPage"`
	IsLastPage        bool  `json:"isLastPage"`
	HasPreviousPage   bool  `json:"hasPreviousPage"`
	HasNextPage       bool  `json:"hasNextPage"`
	NavigatePages     int   `json:"navigatePages"`
	NavigatepageNums  []int `json:"navigatepageNums"`
	NavigateFirstPage int   `json:"navigateFirstPage"`
	NavigateLastPage  int   `json:"navigateLastPage"`
}

type CategoryAndTagReq struct {
	g.Meta `path:"/admin/categoryAndTag" method:"get" summart:"获取标签" tags:"获取标签"`
}

type CategoryAndTagRes struct {
	Categories []Categories `json:"categories"`
	Tags       []Tags       `json:"tags"`
}

type Categories struct {
	Id           int      `json:"id"`
	Categoryname string   `json:"name"`
	Blogs        []string `json:"blogs"`
}
