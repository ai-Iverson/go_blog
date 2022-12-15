package controller

import (
	"context"
	v1 "go_blog/api/v1"
	"go_blog/internal/service"
	"go_blog/utility/utils"
)

var Category = cCategory{}

type cCategory struct {
}

// Category 展示所有标签（分页）
func (c *cCategory) ShowCategory(ctx context.Context, req *v1.ShowCategoryReq) (res *v1.ShowCategoryRes, err error) {
	res = &v1.ShowCategoryRes{}
	allCategory, err := service.Category().ShowCategory(ctx, req.PageNum, req.PageSize)
	utils.MyCopy(ctx, res, allCategory)
	return
}

func (c *cCategory) AddCategory(ctx context.Context, req *v1.AddCategoryReq) (res *v1.AddCategoryRes, err error) {
	err = service.Category().AddCategory(ctx, req.CategoryName)
	return nil, err
}

func (c *cCategory) UpdateCategory(ctx context.Context, req *v1.UpdateCategoryReq) (res *v1.UpdateCategoryRes, err error) {
	err = service.Category().UpdateCategory(ctx, req.Id, req.CategoryName)
	return nil, err
}

func (c *cCategory) DeleteCategory(ctx context.Context, req *v1.DeleteCategoryReq) (res *v1.DeleteCategoryRes, err error) {
	err = service.Category().DeleteCategory(ctx, req.Id)
	return nil, err
}
