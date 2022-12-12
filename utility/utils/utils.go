package utils

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/jinzhu/copier"
	v1 "go_blog/api/v1"
	"go_blog/internal/errorcode"
)

func MyCopy(ctx context.Context, toValue interface{}, fromValue interface{}) (err error) {
	if err = copier.Copy(toValue, fromValue); err != nil {
		return errorcode.NewMyErr(ctx, errorcode.MyInternalError, err)
	} else {
		return nil
	}
}

// Pagination 分页
func Pagination(ctx context.Context, page, size, count, currentSize int) (out *v1.CommonPaginationReq, err error) {
	out = &v1.CommonPaginationReq{}
	r := g.RequestFromCtx(ctx)
	pageInfo := r.GetPage(count, size)

	out.Total = count
	out.PageNum = pageInfo.CurrentPage
	out.Size = currentSize
	out.PageSize = size
	out.StartRow = (page-1)*size + 1
	out.Pages = pageInfo.TotalPage
	out.NavigatePages = 8
	out.NavigateFirstPage = 1

	if pageInfo.TotalPage < 8 {
		out.NavigateLastPage = pageInfo.TotalPage
	} else {
		out.NavigateLastPage = 8
	}
	//判断是否是第一页
	if page == 1 {
		out.IsFirstPage = true
		out.PrePage = 0
		out.HasPreviousPage = false
	} else {
		out.IsFirstPage = false
		out.PrePage = page - 1
		out.HasPreviousPage = true
	}

	// 判断是否是最后一页
	if page == pageInfo.TotalPage {
		out.IsLastPage = true
		out.HasNextPage = false
		out.EndRow = count
		out.NextPage = 0

	} else {
		out.IsLastPage = false
		out.HasNextPage = true
		out.EndRow = page * size
		out.NextPage = page + 1
	}

	for pageNum := 1; pageNum <= pageInfo.TotalPage; pageNum++ {
		out.NavigatepageNums = append(out.NavigatepageNums, pageNum)
	}
	return
}
