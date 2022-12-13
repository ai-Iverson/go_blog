// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
	"go_blog/internal/model"
)

type IBlog interface {
	CreateBlog(ctx context.Context, in model.CreateBlogInput) (err error)
	BlogsList(ctx context.Context, page, size int) (out *model.BlogsListOutput, err error)
	BlogDetail(ctx context.Context, id int) (out *model.BlogDetailOutput, err error)
}

var localBlog IBlog

func Blog() IBlog {
	if localBlog == nil {
		panic("implement not found for interface IBlog, forgot register?")
	}
	return localBlog
}

func RegisterBlog(i IBlog) {
	localBlog = i
}
