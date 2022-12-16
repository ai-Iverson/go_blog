package sitesstings

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"go_blog/internal/dao"
	"go_blog/internal/model"
	"go_blog/internal/service"
)

type sSiteSettings struct {
}

func init() {
	service.RegisterSiteSettings(New())
}

func New() *sSiteSettings {
	return &sSiteSettings{}
}

func (s *sSiteSettings) SiteSettingsDetil(ctx context.Context) (out *model.SiteSettingsDetilOutput, err error) {
	out = &model.SiteSettingsDetilOutput{}
	SiteSettingAll, _ := dao.SiteSetting.Ctx(ctx).All()
	for _, typeValueList := range SiteSettingAll {
		if typeValueList["type"].Int() == 1 {
			TypeStruct1 := model.Type1{}
			gconv.Scan(typeValueList, &TypeStruct1)
			out.Type1 = append(out.Type1, TypeStruct1)

		} else if typeValueList["type"].Int() == 2 {
			TypeStruct2 := model.Type2{}
			gconv.Scan(typeValueList, &TypeStruct2)
			out.Type2 = append(out.Type2, TypeStruct2)

		} else if typeValueList["type"].Int() == 3 {
			TypeStruct3 := model.Type3{}
			gconv.Scan(typeValueList, &TypeStruct3)
			out.Type3 = append(out.Type3, TypeStruct3)

		}
	}

	return
}

func (s *sSiteSettings) UpdateSiteSettings(ctx context.Context, in model.UpdateSiteSettingsInput) (err error) {
	g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		// 删除
		_, err = dao.SiteSetting.Ctx(ctx).WhereIn(dao.SiteSetting.Columns().Id, in.DeleteIds).Delete()
		if err != nil {
			return err
		}
		SiteSettingColumns := dao.SiteSetting.Columns()
		for _, set := range in.Settings {
			glog.Info(ctx, set.ID)
			if set.ID == 0 {
				// 新增
				_, err = dao.SiteSetting.Ctx(ctx).Data(g.Map{
					SiteSettingColumns.Type:   set.Type,
					SiteSettingColumns.Value:  set.Value,
					SiteSettingColumns.NameEn: set.NameEn,
					SiteSettingColumns.NameZh: set.NameZh,
				}).Insert()
			} else {
				// 更新
				_, err = dao.SiteSetting.Ctx(ctx).Data(g.Map{

					SiteSettingColumns.Type:   set.Type,
					SiteSettingColumns.Value:  set.Value,
					SiteSettingColumns.NameEn: set.NameEn,
					SiteSettingColumns.NameZh: set.NameZh,
				}).Where(SiteSettingColumns.Id, set.ID).Update()
			}
		}
		return err
	})
	return err
}
