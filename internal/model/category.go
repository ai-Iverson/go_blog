package model

import v1 "go_blog/api/v1"

type Categories struct {
	Id           int      `json:"id"`
	Categoryname string   `json:"name"`
	Blogs        []string `json:"blogs"`
}

type AllCategoryOutput struct {
	v1.ShowCategoryRes
}
