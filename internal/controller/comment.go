package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
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
	res = &v1.CommentListRes{}
	out, err := service.Comment().CommentList(ctx)
	res.CloseComment = out.CloseComment
	res.AllComment = out.AllComment
	gconv.Scan(out.Comments, &res.Comments)
	return
}
