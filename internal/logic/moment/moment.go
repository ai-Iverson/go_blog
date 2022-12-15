package moment

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"go_blog/internal/dao"
	"go_blog/internal/model"
	"go_blog/internal/model/entity"
	"go_blog/internal/service"
	"go_blog/utility/utils"
)

type sMoment struct {
}

func init() {
	service.RegisterMoment(New())
}

func New() *sMoment {
	return &sMoment{}
}

func (s *sMoment) CreateMoment(ctx context.Context, in model.CreateMomentInput) (err error) {
	m := dao.Moment.Columns()
	_, err = dao.Moment.Ctx(ctx).Data(g.Map{
		m.Content:     in.Content,
		m.IsPublished: in.Published,
		m.Likes:       in.Likes,
		m.CreateTime:  in.CreateTime,
	}).Insert()
	if err != nil {
		return err
	}
	return
}

func (s *sMoment) ShowMoments(ctx context.Context, page, size int) (out *model.ShowMomentOutput, err error) {
	out = &model.ShowMomentOutput{}
	m := dao.Moment.Ctx(ctx).Order("create_time desc")

	count, err := m.Count()
	if err != nil {
		return
	}

	result, err := m.Page(page, size).All()
	if err != nil {
		return
	}

	// 分页信息
	pagination, err := utils.Pagination(ctx, page, size, gconv.Int(count), len(result))
	utils.MyCopy(ctx, &out, pagination)
	gconv.Scan(result, &out.List)

	return
}

func (s *sMoment) UpdatePublishedStatus(ctx context.Context, momentId int, isPublished bool) (err error) {
	_, err = dao.Moment.Ctx(ctx).Data(dao.Moment.Columns().IsPublished, isPublished).
		Where(dao.Moment.Columns().Id, momentId).Update()
	if err != nil {
		return
	}
	return
}

func (s *sMoment) MomoentDetil(ctx context.Context, momentId int) (out *model.MomentDetilOutput, err error) {
	out = &model.MomentDetilOutput{}
	err = dao.Moment.Ctx(ctx).Where(dao.Moment.Columns().Id, momentId).Scan(&out)
	if err != nil {
		return nil, err
	}
	return
}

func (s *sMoment) UpdateMoment(ctx context.Context, in entity.Moment) (err error) {
	m := dao.Moment.Columns()
	_, err = dao.Moment.Ctx(ctx).Data(g.Map{
		m.Likes:       in.Likes,
		m.CreateTime:  in.CreateTime,
		m.IsPublished: in.IsPublished,
		m.Content:     in.Content,
	}).Where(m.Id, in.Id).Update()
	if err != nil {
		return err
	}
	return err
}

func (s *sMoment) DeleteMoment(ctx context.Context, momenId int) (err error) {
	_, err = dao.Moment.Ctx(ctx).Where(dao.Moment.Columns().Id, momenId).Delete()
	if err != nil {
		return err
	}
	return
}
