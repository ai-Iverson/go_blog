package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "go_blog/api/v1"
	"go_blog/internal/service"
)

var Index = cIndex{}

type cIndex struct {
}

func (c *cIndex) CategoryAndTag(ctx context.Context, req *v1.CategoryAndTagReq) (res *v1.CategoryAndTagRes, err error) {
	res = &v1.CategoryAndTagRes{}
	tagsList, err := service.Tags().GetTagsList(ctx)
	if err != nil {
		return nil, err
	}
	for _, tag := range tagsList {
		res.Tags = append(res.Tags, v1.Tags{
			Id:      gconv.Int(tag.Id),
			TagName: tag.TagName,
			Color:   tag.Color,
			Blogs:   []string{},
		})
	}
	categoryList, err := service.Category().GetCategoryList(ctx)
	if err != nil {
		return nil, err
	}
	for _, category := range categoryList {
		res.Categories = append(res.Categories, v1.Categories{
			Id:           gconv.Int(category.Id),
			Categoryname: category.CategoryName,
			Blogs:        []string{},
		})
	}
	return
}
