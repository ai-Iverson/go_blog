package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	v1 "go_blog/api/v1"
	"go_blog/internal/model/entity"
)

type CreateMomentInput struct {
	Content    string      `json:"content"`
	CreateTime *gtime.Time `json:"createTime"`
	Likes      int         `json:"likes"`
	Published  bool        `json:"published"`
}

type ShowMomentOutput struct {
	List []entity.Moment `json:"list"`
	v1.CommonPaginationReq
}

type MomentDetilOutput struct {
	ID         int    `json:"id"`
	Content    string `json:"content"`
	CreateTime string `json:"createTime"`
	Likes      int    `json:"likes"`
	Published  bool   `json:"published"`
}
