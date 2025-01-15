package cmd

import (
	"context"
	"github.com/wangle201210/ai-note/app/ai-note/internal/controller/category"
	"github.com/wangle201210/ai-note/app/ai-note/internal/controller/note"
	"github.com/wangle201210/ai-note/app/ai-note/internal/controller/user"
	"github.com/wangle201210/ai-note/app/ai-note/internal/controller/wechat"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/wangle201210/ai-note/app/ai-note/internal/controller/hello"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.SetWriteTimeout(time.Second * 20)
			s.SetReadTimeout(time.Second * 20)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewV1(),
					user.NewV1(),
					note.NewV1(),
					category.NewV1(),
				)
			})
			s.BindHandler("/wechat_callback", wechat.Server)
			s.Run()
			return nil
		},
	}
)
