package cmd

import (
	"context"
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
				)
			})
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(service.Middleware().TokenAuth)
				//group.Middleware(service.Middleware().ApiAuth)
				group.Bind(
					controller.Dashboart.Dashboart,
				)
			})
			s.Run()
			return nil
		},
	}
)
