package v1

import "github.com/gogf/gf/v2/frame/g"

type AboutReq struct {
	g.Meta        `path:"/admin/about" method:"get" summart:"关于我展示" tags:"关于我展示"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
}

type AboutRes struct {
	MusicID        string `json:"musicId"`
	CommentEnabled string `json:"commentEnabled"`
	Title          string `json:"title"`
	Content        string `json:"content"`
}

type UpdateAboutReq struct {
	g.Meta         `path:"/admin/about" method:"put" summart:"更新关于我" tags:"更新关于我"`
	Authorization  string `json:"Authorization" in:"header"  dc:"Authorization"`
	MusicID        string `json:"musicId"`
	CommentEnabled string `json:"commentEnabled"`
	Title          string `json:"title"`
	Content        string `json:"content"`
}

type UpdateAboutRes struct {
}
