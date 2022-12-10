package tags

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "go_blog/api/v1"
	"go_blog/internal/dao"
	"go_blog/internal/errorcode"
	"go_blog/internal/model"
	"go_blog/internal/model/entity"
	"go_blog/internal/service"
)

type sTags struct {
}

func init() {
	service.RegisterTags(New())
}

func New() *sTags {
	return &sTags{}
}

func (s *sTags) Tags(ctx context.Context, page, size int) (out *model.AllTagsOutput, err error) {
	out = &model.AllTagsOutput{}
	out.List = []v1.Tags{}
	m := dao.Tag.Ctx(ctx).Page(page, size)
	allTags, _ := m.All()
	total, _ := m.Count()
	out.Total = gconv.Int(total)
	err = allTags.Structs(&out.List)
	glog.Info(ctx, "sdf", out)
	return
}

func (s *sTags) AddTags(ctx context.Context, name, color string) (err error) {
	_, err = dao.Tag.Ctx(ctx).Data(entity.Tag{
		Name:  name,
		Color: color,
	}).Insert()
	if err != nil {
		glog.Error(ctx, err)
		return errorcode.NewMyErr(ctx, errorcode.MyInternalError)
	}

	return
}

func (s *sTags) UpdateTags(ctx context.Context, id int, name, color string) (err error) {
	_, err = dao.Tag.Ctx(ctx).Where(dao.Tag.Columns().Id, id).Update(entity.Tag{
		Name:  name,
		Color: color,
	})
	if err != nil {
		glog.Error(ctx, err)
		return errorcode.NewMyErr(ctx, errorcode.MyInternalError)
	}
	return
}

func (s *sTags) DeleteTags(ctx context.Context, id int) (err error) {
	_, err = dao.Tag.Ctx(ctx).Delete(dao.Tag.Columns().Id, id)
	if err != nil {
		glog.Error(ctx, err)
		return errorcode.NewMyErr(ctx, errorcode.MyInternalError)
	}
	return
}
