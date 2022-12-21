package sitesstings

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"go_blog/internal/dao"
	"go_blog/internal/model"
	"go_blog/internal/service"
	"strings"
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

func (s *sSiteSettings) Site(ctx context.Context) (out *model.SiteOutput, err error) {
	out = &model.SiteOutput{}

	// Badges
	Badges, err := s.Badges(ctx)
	if err != nil {
		return nil, err
	}
	gconv.Scan(Badges, &out.Badges)

	//category
	category, err := service.Category().GetCategoryList(ctx)
	if err != nil {
		return nil, err
	}
	gconv.Scan(category, &out.CategoryList)

	// Introduction
	introduction, err := service.SiteSettings().Introduction(ctx)
	if err != nil {
		return nil, err
	}
	gconv.Scan(introduction, &out.Introduction)

	// NewBlogList
	newBlogList, err := service.SiteSettings().NewBlogList(ctx)
	gconv.Scan(newBlogList, &out.NewBlogList)

	//RandomBlogList
	randomBlogList, err := service.SiteSettings().RandomBlogList(ctx)
	gconv.Scan(randomBlogList, &out.RandomBlogList)

	// siteInfo
	siteInfo, err := service.SiteSettings().SiteInfo(ctx)
	gconv.Scan(siteInfo, &out.SiteInfo)

	//taglist
	tagList, err := service.Tags().GetTagsList(ctx)
	for _, tag := range tagList {
		out.TagList = append(out.TagList, model.Tags{
			Id:    gconv.Int(tag.Id),
			Name:  tag.TagName,
			Color: tag.Color,
			Blogs: []string{},
		})
	}
	gconv.Scan(tagList, &out.TagList)

	return
}

func (s *sSiteSettings) Badges(ctx context.Context) (out []*model.Badges, err error) {
	Badge := &model.Badges{}
	siteType3, _ := dao.SiteSetting.Ctx(ctx).Where(dao.SiteSetting.Columns().Type, "3").All()
	for _, bg := range siteType3.List() {
		bgValue := gconv.Map(bg["value"])
		Badge.Title = bgValue["title"].(string)
		Badge.URL = bgValue["url"].(string)
		Badge.Subject = bgValue["subject"].(string)
		Badge.Value = bgValue["value"].(string)
		Badge.Color = bgValue["color"].(string)
		out = append(out, Badge)

	}
	return
}

func (s *sSiteSettings) Introduction(ctx context.Context) (out *model.Introduction, err error) {
	out = &model.Introduction{}
	favorite := &model.Favorites{}
	siteType2, _ := dao.SiteSetting.Ctx(ctx).Where(dao.SiteSetting.Columns().Type, "2").All()
	for _, Int := range siteType2.List() {
		switch Int["name_en"] {
		case "avatar":
			out.Avatar = Int["value"].(string)
		case "name":
			out.Name = Int["value"].(string)
		case "github":
			out.Github = Int["value"].(string)
		case "telegram":
			out.Telegram = Int["value"].(string)
		case "qq":
			out.Qq = Int["value"].(string)
		case "bilibili":
			out.Bilibili = Int["value"].(string)
		case "netease":
			out.Netease = Int["value"].(string)
		case "email":
			out.Email = Int["value"].(string)
		case "rollText":
			rollText := Int["value"].(string)
			rollTextList := strings.Split(rollText, ",")
			gconv.Scan(rollTextList, &out.RollText)
		case "favorite":

			favoriteValue := gconv.Map(Int["value"])
			//gconv.Scan(favoriteValue, favorite)
			favorite.Title = favoriteValue["title"].(string)
			favorite.Content = favoriteValue["content"].(string)
			out.Favorites = append(out.Favorites, *favorite)

		}
	}
	return
}

func (s *sSiteSettings) NewBlogList(ctx context.Context) (out []*model.NewBlogList, err error) {
	err = dao.Blog.Ctx(ctx).OrderDesc(dao.Blog.Columns().CreateTime).Limit(3).Scan(&out)
	return
}

func (s *sSiteSettings) RandomBlogList(ctx context.Context) (out []*model.RandomBlogList, err error) {
	err = dao.Blog.Ctx(ctx).OrderRandom().Limit(5).Scan(&out)
	return

}

func (s *sSiteSettings) SiteInfo(ctx context.Context) (out *model.SiteInfo, err error) {
	out = &model.SiteInfo{}
	copyRight := &model.Copyright{}
	siteType1, _ := dao.SiteSetting.Ctx(ctx).Where(dao.SiteSetting.Columns().Type, "1").All()
	for _, st := range siteType1.List() {
		switch st["name_en"] {
		case "blogName":
			out.BlogName = st["value"].(string)
		case "webTitleSuffix":
			out.WebTitleSuffix = st["value"].(string)

		case "footerImgTitle":
			out.FooterImgTitle = st["value"].(string)
		case "footerImgUrl":
			out.FooterImgURL = st["value"].(string)
		case "beian":
			out.Beian = st["value"].(string)
		case "reward":
			out.Reward = st["value"].(string)
		case "commentAdminFlag":
			out.CommentAdminFlag = st["value"].(string)
		case "copyright":
			copyRightValue := gconv.Map(st["value"])
			copyRight.Title = copyRightValue["title"].(string)
			copyRight.SiteName = copyRightValue["siteName"].(string)
			gconv.Scan(copyRight, &out.Copyright)
		}
	}
	return
}
