package blog

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "go_blog/api/v1"
	"go_blog/internal/dao"
	"go_blog/internal/model"
	"go_blog/internal/model/entity"
	"go_blog/internal/service"
	"go_blog/utility/utils"
)

type sBlog struct {
}

func init() {
	service.RegisterBlog(New())
}

func New() *sBlog {
	return &sBlog{}
}

func (s sBlog) CreateBlog(ctx context.Context, in model.CreateBlogInput) (err error) {
	// 数据库事务操作
	g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		blogId, err := tx.Ctx(ctx).InsertAndGetId("blog", entity.Blog{
			Title:            in.Title,
			FirstPicture:     in.FirstPicture,
			Content:          in.Content,
			Description:      in.Description,
			IsPublished:      gconv.Int(in.IsPublished),
			IsRecommend:      gconv.Int(in.IsRecommend),
			IsAppreciation:   gconv.Int(in.IsAppreciation),
			IsCommentEnabled: gconv.Int(in.IsCommentEnabled),
			CreateTime:       gtime.Now(),
			UpdateTime:       gtime.Now(),
			Views:            in.Views,
			Words:            gconv.Int(in.Words),
			ReadTime:         in.ReadTime,
			CategoryId:       gconv.Int64(in.CategoryId),
			IsTop:            gconv.Int(in.IsTop),
			Password:         in.Password,
			UserId:           gconv.Int64(in.User),
		})
		if err != nil {
			return err
		}
		for _, tag := range in.TagList {
			_, err = dao.BlogTag.Ctx(ctx).Data(entity.BlogTag{
				BlogId: blogId,
				TagId:  gconv.Int64(tag),
			}).Insert()
		}

		return err
	})
	return
}

func (s sBlog) UpdateBlog(ctx context.Context, in v1.UpdateBlogReq) (err error) {
	g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		blogCls := dao.Blog.Columns()
		_, err := tx.Ctx(ctx).Update(dao.Blog.Table(), g.Map{
			blogCls.Title:            in.Title,
			blogCls.FirstPicture:     in.FirstPicture,
			blogCls.Content:          in.Content,
			blogCls.Description:      in.Description,
			blogCls.IsPublished:      gconv.Int(in.Published),
			blogCls.IsRecommend:      gconv.Int(in.Recommend),
			blogCls.IsAppreciation:   gconv.Int(in.Appreciation),
			blogCls.IsCommentEnabled: gconv.Int(in.CommentEnabled),
			blogCls.Views:            in.Views,
			blogCls.Words:            gconv.Int(in.Words),
			blogCls.ReadTime:         in.ReadTime,
			blogCls.CategoryId:       gconv.Int64(in.Cate),
			blogCls.IsTop:            gconv.Int(in.Top),
			blogCls.Password:         in.Password,
			blogCls.UpdateTime:       gtime.Now(),
		}, dao.Blog.Columns().Id, in.Id)
		if err != nil {
			return err
		}

		//删除关联标签
		_, err = dao.BlogTag.Ctx(ctx).Where(dao.BlogTag.Columns().BlogId, in.Id).Delete()
		glog.Info(ctx, "gconv.Int64(in.Id", gconv.Int64(in.Id))
		// 关联标签
		for _, tag := range in.TagList {
			_, err = dao.BlogTag.Ctx(ctx).Data(entity.BlogTag{
				BlogId: gconv.Int64(in.Id),
				TagId:  gconv.Int64(tag),
			}).Insert()
			if err != nil {
				return err
			}
		}

		return err
	})
	return
}

func (s *sBlog) BlogsList(ctx context.Context, title string, categroyId, page, size int) (out *model.BlogsListOutput, err error) {
	out = &model.BlogsListOutput{}
	b := dao.Blog.Ctx(ctx).Order("is_top desc,create_time desc")

	// 标题查询
	if len(title) > 0 {
		b = b.WhereLike(dao.Blog.Columns().Title, "%"+title+"%")
	}

	// 分类查询
	if categroyId != 0 {
		b = b.Where(dao.Blog.Columns().CategoryId, categroyId)
	}

	// 总数
	count, err := b.Count()
	if err != nil {
		return nil, err
	}

	result, err := b.Page(page, size).All()
	if err != nil {
		return nil, err
	}
	// 分页信息
	pagination, err := utils.Pagination(ctx, page, size, gconv.Int(count), len(result))
	utils.MyCopy(ctx, &out, pagination)

	err = gconv.Scan(result, &out.List)
	if err != nil {
		return
	}

	// 查询每个blog的categroy信息
	err = dao.Category.Ctx(ctx).
		Where(dao.Category.Columns().Id, gdb.ListItemValuesUnique(out.List, "CategoryId")).
		ScanList(&out.List, "Category", "id:CategoryId")
	return

}

func (s *sBlog) BlogDetail(ctx context.Context, id int) (out *model.BlogDetailOutput, err error) {
	out = &model.BlogDetailOutput{}
	err = dao.Blog.Ctx(ctx).Where(dao.Blog.Columns().Id, id).Scan(&out)
	if err != nil {
		return
	}

	err = dao.Category.Ctx(ctx).Where(dao.Category.Columns().Id, out.CategoryId).Scan(&out.Category)
	if err != nil {
		return
	}

	tagList := []*entity.BlogTag{}
	err = dao.BlogTag.Ctx(ctx).Where(dao.BlogTag.Columns().BlogId, id).Scan(&tagList)
	if err != nil {
		return
	}

	err = dao.Tag.Ctx(ctx).Where(dao.Tag.Columns().Id, gdb.ListItemValuesUnique(tagList, "TagId")).Scan(&out.Tags)
	if err != nil {
		return
	}

	return

}

func (s *sBlog) UpdateBlogTop(ctx context.Context, blogId int, top bool) (err error) {
	res, err := dao.Blog.Ctx(ctx).Data(dao.Blog.Columns().IsTop, top).Where(dao.Blog.Columns().Id, blogId).Update()
	if res == nil {
		glog.Error(ctx, "数据更新错误")
	}
	return
}

func (s *sBlog) UpdateBlogRecommend(ctx context.Context, blogId int, recommend bool) (err error) {
	res, err := dao.Blog.Ctx(ctx).Data(dao.Blog.Columns().IsRecommend, recommend).Where(dao.Blog.Columns().Id, blogId).Update()
	if res == nil {
		glog.Error(ctx, "数据更新错误")
	}
	return
}

func (s *sBlog) UpdateBlogVisibility(ctx context.Context, in model.UpdateBlogVisibilityInput) (err error) {
	blogCls := dao.Blog.Columns()
	res, err := dao.Blog.Ctx(ctx).Data(g.Map{
		blogCls.IsTop:            in.Top,
		blogCls.IsRecommend:      in.Recommend,
		blogCls.IsPublished:      in.Published,
		blogCls.IsAppreciation:   in.Appreciation,
		blogCls.Password:         in.Password,
		blogCls.IsCommentEnabled: in.CommentEnabled,
	}).Where(dao.Blog.Columns().Id, in.Id).Update()
	if res == nil {
		glog.Error(ctx, "数据更新错误")
	}
	return
}

func (s *sBlog) DeleteBlog(ctx context.Context, blogId int) (err error) {
	g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.Blog.Ctx(ctx).Where(dao.Blog.Columns().Id, blogId).Delete()
		if err != nil {
			return err
		}
		//删除关联标签
		_, err = dao.BlogTag.Ctx(ctx).Where(dao.BlogTag.Columns().BlogId, blogId).Delete()
		if err != nil {
			return err
		}
		return err
	})
	return err
}
