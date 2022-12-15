package tags

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"go_blog/internal/dao"
	"go_blog/internal/errorcode"
	"go_blog/internal/model"
	"go_blog/internal/model/entity"
	"go_blog/internal/service"
	"go_blog/utility/utils"
)

type sTags struct {
}

func init() {
	service.RegisterTags(New())
}

func New() *sTags {
	return &sTags{}
}

func (s *sTags) GetTagsList(ctx context.Context) (out []*entity.Tag, err error) {
	err = dao.Tag.Ctx(ctx).Scan(&out)
	return
}

func (s *sTags) ShowTags(ctx context.Context, page, size int) (out *model.AllTagsOutput, err error) {
	out = &model.AllTagsOutput{}
	t := dao.Tag.Ctx(ctx)
	// 总数
	count, err := t.Count()
	if err != nil {
		return nil, err
	}

	result, err := dao.Tag.Ctx(ctx).Page(page, size).All()
	if err != nil {
		return nil, err
	}

	pagination, err := utils.Pagination(ctx, page, size, gconv.Int(count), len(result))
	if err != nil {
		return nil, err
	}
	utils.MyCopy(ctx, &out, pagination)
	err = gconv.Structs(result, &out.List)

	return
}

func (s *sTags) AddTags(ctx context.Context, name, color string) (err error) {
	_, err = dao.Tag.Ctx(ctx).Data(entity.Tag{
		TagName: name,
		Color:   color,
	}).Insert()
	if err != nil {
		glog.Error(ctx, err)
		return errorcode.NewMyErr(ctx, errorcode.MyInternalError)
	}

	return
}

func (s *sTags) UpdateTags(ctx context.Context, id int, name, color string) (err error) {
	_, err = dao.Tag.Ctx(ctx).Where(dao.Tag.Columns().Id, id).Update(entity.Tag{
		TagName: name,
		Color:   color,
	})
	if err != nil {
		glog.Error(ctx, err)
		return errorcode.NewMyErr(ctx, errorcode.MyInternalError)
	}
	return
}

func (s *sTags) DeleteTags(ctx context.Context, id int) (err error) {
	blogTag, err := dao.BlogTag.Ctx(ctx).Where(dao.BlogTag.Columns().TagId, id).All()
	if blogTag.Len() == 0 {
		_, err = dao.Tag.Ctx(ctx).Delete(dao.Tag.Columns().Id, id)
		if err != nil {
			glog.Error(ctx, err)
			return errorcode.NewMyErr(ctx, errorcode.MyInternalError)
		}
	} else {
		return errorcode.NewMyErr(ctx, errorcode.BadDelete, "已有Blog关联此Tag")
	}

	return
}
