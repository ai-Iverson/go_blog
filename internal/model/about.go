package model

type AboutOutput struct {
	MusicID        string `json:"musicId"`
	CommentEnabled string `json:"commentEnabled"`
	Title          string `json:"title"`
	Content        string `json:"content"`
}
