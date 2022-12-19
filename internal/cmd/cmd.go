package cmd

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"go_blog/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"go_blog/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 代替默认的log
			SetLoggerDefaultHandler()
			s := g.Server()
			s.Use(
				service.Middleware().MiddlewareCORS,
				service.Middleware().Ctx,
				service.Middleware().I18NMiddleware,
				service.Middleware().ResponseHandler,
			)

			s.BindStatusHandler(404, func(r *ghttp.Request) {
				r.Response.Writeln("404 - 你似乎来到了没有知识存在的荒原…")
			})
			s.BindStatusHandler(403, func(r *ghttp.Request) {
				r.Response.Writeln("403 - 网站拒绝显示此网页")
			})

			//// OpenApi自定义信息
			//oai.Info.Title = `API Reference`
			//oai.Config.CommonResponse = response.JsonRes{}
			//oai.Config.CommonResponseDataField = `Data`

			s.Group("/", func(group *ghttp.RouterGroup) {
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.Hello,
					controller.User.Login,
					controller.Index.CategoryAndTag,
				)
			})
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(service.Middleware().TokenAuth)
				//group.Middleware(service.Middleware().ApiAuth)
				group.Bind(
					controller.Dashboart.Dashboart,
					controller.Tags.ShowTags,
					controller.Tags.AddTags,
					controller.Tags.UpdateTags,
					controller.Tags.DeleteTags,
					controller.Category.ShowCategory,
					controller.Category.AddCategory,
					controller.Category.UpdateCategory,
					controller.Category.DeleteCategory,
					controller.Blog.CreateBlog,
					controller.Blog.ShowBlogs,
					controller.Blog.UpdateBlog,
					controller.Blog.DeleteBlog,
					controller.Blog.BlogDetail,
					controller.Blog.UpdateBlogTop,
					controller.Blog.UpdateBlogRecommend,
					controller.Blog.UpdateBlogVisibilityReq,
					controller.Moment.CteatMomet,
					controller.Moment.ShoewMoments,
					controller.Moment.UpdatePublishedStatus,
					controller.Moment.MomoentDetil,
					controller.Moment.UpdateMoment,
					controller.Moment.DeleteMoment,
					controller.SiteSttings.SiteSettingsDetil,  // 站点设置
					controller.SiteSttings.UpdateSiteSettings, // 更新站点设置
					controller.About.About,                    // 关于我
					controller.About.UpdateAbout,              // 更新关于我
				)
			})
			s.Run()
			return nil
		},
	}
)

// SetLoggerDefaultHandler 替代默认的日志handler，禁止控制台输出，全部输出到文件
func SetLoggerDefaultHandler() {
	glog.SetDefaultHandler(func(ctx context.Context, in *glog.HandlerInput) {
		m := map[string]interface{}{
			"stdout":            g.Config().MustGet(ctx, "logger.globalStdout", true).Bool(), // 使用自定义的全局字段
			"path":              g.Config().MustGet(ctx, "logger.path", "log/").String(),     // 此处必须重新设置，才可以实现db的log写入文件
			"writerColorEnable": true,
		}

		in.Logger.SetConfigWithMap(m)
		in.Next(ctx)
	})
	// 添加日志代码行号和年月日
	glog.SetFlags(glog.F_TIME_STD | glog.F_FILE_SHORT)
}
