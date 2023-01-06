package v1

import "github.com/gogf/gf/v2/frame/g"

type CommentReq struct {
	g.Meta          `path:"/comment" method:"post" summart:"评论" tags:"评论"`
	Content         string `json:"content"`
	Nickname        string `json:"nickname"`
	Email           string `json:"email"`
	Website         string `json:"website"`
	Notice          bool   `json:"notice"`
	Page            int    `json:"page"`
	BlogID          int    `json:"blogId"`
	ParentCommentID int    `json:"parentCommentId"`
}

type CommentRes struct {
}

type CommentListReq struct {
	g.Meta   `path:"/comments" method:"get" summart:"评论列表" tags:"评论列表"`
	Page     int `json:"page"`
	BlogId   int `json:"blogId"`
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

type CommentListRes struct {
	Comments     Comments `json:"comments"`
	AllComment   int      `json:"allComment"`
	CloseComment int      `json:"closeComment"`
}
type List struct {
	ID                    int           `json:"id"`
	Nickname              string        `json:"nickname"`
	Content               string        `json:"content"`
	Avatar                string        `json:"avatar"`
	CreateTime            string        `json:"createTime"`
	Website               string        `json:"website"`
	AdminComment          bool          `json:"adminComment"`
	ParentCommentID       string        `json:"parentCommentId"`
	ParentCommentNickname interface{}   `json:"parentCommentNickname"`
	ReplyComments         []interface{} `json:"replyComments"`
}
type Comments struct {
	TotalPage int    `json:"totalPage"`
	List      []List `json:"list"`
}
