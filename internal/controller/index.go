package controller

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	v1 "go_blog/api/v1"
	"go_blog/internal/service"
	"go_blog/utility/utils"
)

var Index = cIndex{}

type cIndex struct {
}

func (c *cIndex) CategoryAndTag(ctx context.Context, req *v1.CategoryAndTagReq) (res *v1.CategoryAndTagRes, err error) {
	res = &v1.CategoryAndTagRes{}
	out, err := service.Index().CategoryAndTag(ctx)
	glog.Info(ctx, "123", out)
	utils.MyCopy(ctx, res, out)
	glog.Info(ctx, "xxxx", res)
	return
}

func (c *cIndex) Tags(ctx context.Context, req *v1.TagsReq) (res *v1.TagsRes, err error) {
	res = &v1.TagsRes{}
	allTags, err := service.Tags().Tags(ctx, req.PageNum, req.PageSize)
	utils.MyCopy(ctx, res, allTags)
	glog.Info(ctx, "eeww", res)
	return
}
