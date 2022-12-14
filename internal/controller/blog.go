package controller

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "go_blog/api/v1"
	"go_blog/internal/model"
	"go_blog/internal/service"
)

var Blog = cBlog{}

type cBlog struct {
}

// CreateBlog 写文章
func (c cBlog) CreateBlog(ctx context.Context, req *v1.CreateBlogReq) (res *v1.CreateBlogRes, err error) {
	userToken := service.Context().Get(ctx)
	userId := userToken.Token.UserKey
	glog.Info(ctx, "userID", userToken, userId, gconv.Int64(userId))
	err = service.Blog().CreateBlog(ctx, model.CreateBlogInput{
		Title:            req.Title,
		FirstPicture:     req.FirstPicture,
		Content:          req.Content,
		Description:      req.Description,
		IsPublished:      req.Published,
		IsRecommend:      req.Recommend,
		IsAppreciation:   req.Appreciation,
		IsCommentEnabled: req.CommentEnabled,
		Views:            req.Views,
		Words:            req.Words,
		ReadTime:         req.ReadTime,
		CategoryId:       req.Cate,
		IsTop:            req.Top,
		Password:         req.Password,
		User:             userId,
		TagList:          req.TagList,
	})
	return
}

// ShowBlogs 展示文章
func (c cBlog) ShowBlogs(ctx context.Context, req *v1.ShowBlogsReq) (res *v1.ShowBlogsRes, err error) {
	res = &v1.ShowBlogsRes{}

	// 文章列表
	blogList, err := service.Blog().BlogsList(ctx, req.Title, req.CategoryId, req.PageNum, req.PageSize)
	err = gconv.Scan(blogList, &res.Blogs)
	if err != nil {
		return
	}

	// 分类列表
	categoryList, err := service.Category().GetCategoryList(ctx)
	glog.Info(ctx, categoryList)
	if err != nil {
		return nil, err
	}

	// 所有categroy
	for _, category := range categoryList {
		res.Categories = append(res.Categories, v1.Categories{
			Id:           gconv.Int(category.Id),
			Categoryname: category.CategoryName,
			Blogs:        []string{},
		})
	}

	// 添加blog []
	for i := 0; i < len(res.Blogs.List); i++ {
		res.Blogs.List[i].Category.Blogs = []string{}
	}
	return
}

// 文章详情
func (c cBlog) BlogDetail(ctx context.Context, req *v1.BlogDetailReq) (res *v1.BlogDetailRes, err error) {
	res = &v1.BlogDetailRes{}

	out, err := service.Blog().BlogDetail(ctx, req.Id)
	gconv.Scan(out, &res)
	//utils.MyCopy(ctx, res, out)

	return
}

//  修改文章置顶信息
func (c *cBlog) UpdateBlogTop(ctx context.Context, req *v1.UpdateBlogTopReq) (res *v1.UpdateBlogTopRes, err error) {
	err = service.Blog().UpdateBlogTop(ctx, req.Id, req.Top)
	return
}
