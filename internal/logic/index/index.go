package index

import (
	"context"
	"go_blog/internal/dao"
	"go_blog/internal/model"
	"go_blog/internal/service"
)

type sIndex struct {
}

func init() {
	service.RegisterIndex(New())
}

func New() *sIndex {
	return &sIndex{}
}

func (s *sIndex) CategoryAndTag(ctx context.Context) (out *model.CategoryAndTagOutput, err error) {

	out = &model.CategoryAndTagOutput{}
	categoriesList, err := dao.Category.Ctx(ctx).All()

	out.Categories = []model.Categories{}
	err = categoriesList.Structs(&out.Categories)

	tagsList, err := dao.Tag.Ctx(ctx).All()
	err = tagsList.Structs(&out.Tags)

	return out, err

}
