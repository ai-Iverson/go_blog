package controller

import (
	"context"
	v1 "go_blog/api/v1"
	"go_blog/internal/model"
	"go_blog/internal/service"
	"go_blog/utility/utils"
)

var Comment = cComment{}

type cComment struct {
}

// Comment 发表评论
func (c *cComment) Comment(ctx context.Context, req *v1.CommentReq) (res *v1.CommentRes, err error) {
	in := &model.CommentInput{}
	utils.MyCopy(ctx, in, req)
	service.Comment().Comment(ctx, *in)
	return
}

// CommentList 评论列表
func (c *cComment) CommentList(ctx context.Context, req *v1.CommentListReq) (res *v1.CommentListRes, err error) {
	service.Comment().CommentList(ctx)
	return
}
