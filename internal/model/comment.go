package model

type CommentInput struct {
	Content         string `json:"content"`
	Nickname        string `json:"nickname"`
	Email           string `json:"email"`
	Website         string `json:"website"`
	Notice          bool   `json:"notice"`
	Page            int    `json:"page"`
	BlogID          int    `json:"blogId"`
	ParentCommentID int    `json:"parentCommentId"`
}

type CommentListOutput struct {
	Comments     Comments `json:"comments"`
	AllComment   int      `json:"allComment"`
	CloseComment int      `json:"closeComment"`
}
type List struct {
	ID                    int    `json:"id"`
	Nickname              string `json:"nickname"`
	Content               string `json:"content"`
	Avatar                string `json:"avatar"`
	CreateTime            string `json:"createTime"`
	Website               string `json:"website"`
	AdminComment          bool   `json:"adminComment"`
	ParentCommentID       string `json:"parentCommentId"`
	ParentCommentNickname string `json:"parentCommentNickname"`
	ReplyComments         []List `json:"replyComments"`
}
type Comments struct {
	TotalPage int    `json:"totalPage"`
	List      []List `json:"list"`
}
