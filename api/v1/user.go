package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/admin/login" method:"post" summart:"用户登录" tags:"登录"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
type LoginRes struct {
	User  *UserGetInfoRes `json:"user"`
	Token string          `json:"token"`
}

type UserGetInfoRes struct {
	Id         int         `json:"id"`
	Username   string      `json:"username"`
	Password   interface{} `json:"password"`
	Nickname   string      `json:"nickname"`
	Avatar     string      `json:"avatar"`
	Email      string      `json:"email"`
	CreateTime int64       `json:"createTime"`
	UpdateTime int64       `json:"updateTime"`
	Role       string      `json:"role"`
}
