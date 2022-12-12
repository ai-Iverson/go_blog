package controller

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	v1 "go_blog/api/v1"
	"go_blog/internal/service"
	"go_blog/utility/utils"
)

var Tags = cTags{}

type cTags struct {
}

// Tags 展示所有标签（分页）
func (c *cTags) ShowTags(ctx context.Context, req *v1.ShowTagsReq) (res *v1.ShowTagsRes, err error) {
	res = &v1.ShowTagsRes{}
	allTags, err := service.Tags().ShowTags(ctx, req.PageNum, req.PageSize)
	utils.MyCopy(ctx, res, allTags)
	glog.Info(ctx, "eeww", res)
	return
}

func (c *cTags) AddTags(ctx context.Context, req *v1.AddTagsReq) (res *v1.AddTagsRes, err error) {
	err = service.Tags().AddTags(ctx, req.TagName, req.Color)
	return nil, err
}

func (c *cTags) UpdateTags(ctx context.Context, req *v1.UpdateTagsReq) (res *v1.UpdateTagsRes, err error) {
	err = service.Tags().UpdateTags(ctx, req.Id, req.TagName, req.Color)
	return nil, err
}

func (c *cTags) DeleteTags(ctx context.Context, req *v1.DeleteTagsReq) (res *v1.DeleteTagsRes, err error) {
	err = service.Tags().DeleteTags(ctx, req.Id)
	return nil, err
}
