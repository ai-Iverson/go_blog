package about

import (
	"context"
	"go_blog/internal/dao"
	"go_blog/internal/model"
	"go_blog/internal/service"
)

type sAbout struct {
}

func init() {
	service.RegisterAbout(New())
}

func New() *sAbout {
	return &sAbout{}
}

func (s *sAbout) About(ctx context.Context) (out *model.AboutOutput, err error) {
	out = &model.AboutOutput{}
	//abouts := &entity.About{}
	about, err := dao.About.Ctx(ctx).All()
	if err != nil {
		return nil, err
	}
	for _, ab := range about {
		switch ab["name_en"].String() {
		case "title":
			out.Title = ab["value"].String()
		case "musicId":
			out.MusicID = ab["value"].String()
		case "content":
			out.Content = ab["value"].String()
		case "commentEnabled":
			out.CommentEnabled = ab["value"].String()

		}
	}
	return
}

func (s *sAbout) UpdateAbout(ctx context.Context, in model.AboutOutput) (err error) {
	dao.About.Ctx(ctx).Data(dao.About.Columns().Value, in.Title).Where(dao.About.Columns().NameEn, "title").Update()
	dao.About.Ctx(ctx).Data(dao.About.Columns().Value, in.MusicID).Where(dao.About.Columns().NameEn, "musicId").Update()
	dao.About.Ctx(ctx).Data(dao.About.Columns().Value, in.Content).Where(dao.About.Columns().NameEn, "Content").Update()
	dao.About.Ctx(ctx).Data(dao.About.Columns().Value, in.CommentEnabled).Where(dao.About.Columns().NameEn, "CommentEnabled").Update()
	return err
}
