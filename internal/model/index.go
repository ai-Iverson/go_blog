package model

import v1 "go_blog/api/v1"

type CategoryAndTagOutput struct {
	Categories []Categories `json:"categories"`
	Tags       []Tags       `json:"tags"`
}

type Categories struct {
	Id    int      `json:"id"`
	Name  string   `json:"name"`
	Blogs []string `json:"blogs"`
}

type Tags struct {
	Id    int      `json:"id"`
	Name  string   `json:"name"`
	Color string   `json:"color"`
	Blogs []string `json:"blogs"`
}

type AllTagsOutput struct {
	v1.TagsRes
}
