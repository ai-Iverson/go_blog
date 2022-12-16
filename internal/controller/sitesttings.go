package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "go_blog/api/v1"
	"go_blog/internal/model"
	"go_blog/internal/service"
	"go_blog/utility/utils"
)

var SiteSttings = cSiteSttings{}

type cSiteSttings struct {
}

// SiteSettingsDetil 站点设置
func (c *cSiteSttings) SiteSettingsDetil(ctx context.Context, req *v1.SiteSettingsDetilReq) (res *v1.SiteSettingsDetilRes, err error) {
	res = &v1.SiteSettingsDetilRes{}
	out, err := service.SiteSettings().SiteSettingsDetil(ctx)
	if err != nil {
		return nil, err
	}
	gconv.Scan(out, &res)
	return

}

// UpdateSiteSettingsRes 更新站点设置
func (c *cSiteSttings) UpdateSiteSettings(ctx context.Context, req *v1.UpdateSiteSettingsReq) (res *v1.UpdateSiteSettingsRes, err error) {
	in := &model.UpdateSiteSettingsInput{}
	utils.MyCopy(ctx, &in, req)
	service.SiteSettings().UpdateSiteSettings(ctx, *in)

	return nil, err
}
