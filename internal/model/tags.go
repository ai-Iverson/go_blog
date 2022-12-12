package model

import v1 "go_blog/api/v1"

type Tags struct {
	Id    int      `json:"id"`
	Name  string   `json:"name"`
	Color string   `json:"color"`
	Blogs []string `json:"blogs"`
}

type AllTagsOutput struct {
	v1.ShowTagsRes
}
