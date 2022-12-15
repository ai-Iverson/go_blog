package controller

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "go_blog/api/v1"
	"go_blog/internal/model"
	"go_blog/internal/model/entity"
	"go_blog/internal/service"
	"go_blog/utility/utils"
)

var Moment = cMoment{}

type cMoment struct {
}

// CteatMomet 创建动态
func (c *cMoment) CteatMomet(ctx context.Context, req *v1.CreateMomentReq) (res *v1.CreateMomentRes, err error) {
	err = service.Moment().CreateMoment(ctx, model.CreateMomentInput{
		Content:    req.Content,
		CreateTime: req.CreateTime,
		Likes:      req.Likes,
		Published:  req.Published,
	})
	if err != nil {
		return nil, err
	}
	return
}

// ShoewMoments 展示动态
func (c *cMoment) ShoewMoments(ctx context.Context, req *v1.ShowMomentsReq) (res *v1.ShowMomentsRes, err error) {

	res = &v1.ShowMomentsRes{}
	out, err := service.Moment().ShowMoments(ctx, req.PageNum, req.PageSize)
	glog.Info(ctx, "out", out)

	utils.MyCopy(ctx, &res, out)

	return
}

// UpdatePublishedStatus 修改动态发布状态
func (c *cMoment) UpdatePublishedStatus(ctx context.Context, req *v1.UpdatePublishedStatusReq) (res *v1.UpdatePublishedStatusRes, err error) {

	err = service.Moment().UpdatePublishedStatus(ctx, req.Id, req.Published)
	if err != nil {
		return nil, err
	}
	return
}

// MomoentDetil 动态详情
func (c *cMoment) MomoentDetil(ctx context.Context, req *v1.MomoentDetilReq) (res *v1.MomoentDetilRes, err error) {
	res = &v1.MomoentDetilRes{}
	out, err := service.Moment().MomoentDetil(ctx, req.Id)
	gconv.Scan(out, &res)
	return
}

// UpdateMoment 更新动态
func (c *cMoment) UpdateMoment(ctx context.Context, req *v1.UpdateMomentReq) (res *v1.UpdateMomentRes, err error) {
	err = service.Moment().UpdateMoment(ctx, entity.Moment{
		Id:          gconv.Int64(req.Id),
		Content:     req.Content,
		CreateTime:  req.CreateTime,
		Likes:       req.Likes,
		IsPublished: gconv.Int(req.Published),
	})
	return
}

func (c *cMoment) DeleteMoment(ctx context.Context, req *v1.DeleteMomentReq) (res *v1.DeleteMomentRes, err error) {
	err = service.Moment().DeleteMoment(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return
}
