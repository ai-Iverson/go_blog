package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type Moment struct {
	Content    string      `json:"content"`
	CreateTime *gtime.Time `json:"createTime"`
	Likes      int         `json:"likes"`
	Published  bool        `json:"published"`
}

type CreateMomentReq struct {
	g.Meta        `path:"/admin/moment" method:"post" summart:" 创建动态" tags:"创建动态"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Moment
}

type CreateMomentRes struct {
}

type ShowMomentsReq struct {
	g.Meta        `path:"/admin/moments" method:"get" summart:" 展示动态" tags:"展示动态"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	PageNum       int    `json:"pageNum"`
	PageSize      int    `json:"pageSize"`
}

type MomentsList struct {
	Id int `json:"id"`
	Moment
}

type ShowMomentsRes struct {
	List []MomentsList `json:"list"`
	CommonPaginationReq
}

type UpdatePublishedStatusReq struct {
	g.Meta        `path:"/admin/moment/published" method:"put" summart:" 修改动态发布状态" tags:"修改动态发布状态"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Id            int    `json:"Id"`
	Published     bool   `json:"published"`
}

type UpdatePublishedStatusRes struct {
}

type MomoentDetilReq struct {
	g.Meta        `path:"/admin/moment" method:"get" summart:" 动态详情" tags:"动态详情"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Id            int    `json:"id"`
}

type MomoentDetilRes struct {
	ID         int    `json:"id"`
	Content    string `json:"content"`
	CreateTime string `json:"createTime"`
	Likes      int    `json:"likes"`
	Published  bool   `json:"published"`
}

type UpdateMomentReq struct {
	g.Meta        `path:"/admin/moment" method:"put" summart:" 更新动态" tags:"更新动态"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Id            int    `json:"id"`
	Moment
}

type UpdateMomentRes struct {
}

type DeleteMomentReq struct {
	g.Meta        `path:"/admin/moment" method:"delete" summart:" 删除动态" tags:"删除动态"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Id            int    `json:"id"`
}

type DeleteMomentRes struct {
}
