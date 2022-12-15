package category

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

type sCategory struct {
}

func init() {
	service.RegisterCategory(New())
}

func New() *sCategory {
	return &sCategory{}
}

func (s *sCategory) GetCategoryList(ctx context.Context) (out []*entity.Category, err error) {
	err = dao.Category.Ctx(ctx).Scan(&out)
	return
}

func (s *sCategory) ShowCategory(ctx context.Context, page, size int) (out *model.AllCategoryOutput, err error) {
	out = &model.AllCategoryOutput{}
	t := dao.Category.Ctx(ctx)
	// 总数
	count, err := t.Count()
	if err != nil {
		return nil, err
	}

	result, err := t.Page(page, size).All()
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

func (s *sCategory) AddCategory(ctx context.Context, name string) (err error) {
	_, err = dao.Category.Ctx(ctx).Data(entity.Category{
		CategoryName: name,
	}).Insert()
	if err != nil {
		glog.Error(ctx, err)
		return errorcode.NewMyErr(ctx, errorcode.MyInternalError)
	}

	return
}

func (s *sCategory) UpdateCategory(ctx context.Context, id int, name string) (err error) {
	_, err = dao.Category.Ctx(ctx).Where(dao.Category.Columns().Id, id).Update(entity.Category{
		CategoryName: name,
	})
	if err != nil {
		glog.Error(ctx, err)
		return errorcode.NewMyErr(ctx, errorcode.MyInternalError)
	}
	return
}

func (s *sCategory) DeleteCategory(ctx context.Context, id int) (err error) {
	categoryBlog, err := dao.Blog.Ctx(ctx).Where(dao.Blog.Columns().CategoryId, id).All()
	if categoryBlog.Len() == 0 {
		_, err = dao.Category.Ctx(ctx).Delete(dao.Category.Columns().Id, id)
		if err != nil {
			glog.Error(ctx, err)
			return errorcode.NewMyErr(ctx, errorcode.MyInternalError)
		}
	}

	return errorcode.NewMyErr(ctx, errorcode.BadDelete, "已有blog关联此categroy")

}
