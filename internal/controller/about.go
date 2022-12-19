package controller

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "go_blog/api/v1"
	"go_blog/internal/model"
	"go_blog/internal/service"
)

var About = cAbout{}

type cAbout struct {
}

// About 关于自己
func (c *cAbout) About(ctx context.Context, req *v1.AboutReq) (res *v1.AboutRes, err error) {
	about, err := service.About().About(ctx)
	if err != nil {
		return nil, err
	}
	gconv.Struct(about, &res)
	return
}

// 更新关于自己

func (c *cAbout) UpdateAbout(ctx context.Context, req *v1.UpdateAboutReq) (res *v1.UpdateAboutRes, err error) {
	in := &model.AboutOutput{}
	gconv.Struct(req, &in)
	glog.Info(ctx, in)
	err = service.About().UpdateAbout(ctx, *in)
	return nil, err
}
