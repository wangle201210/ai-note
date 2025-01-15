package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/wangle201210/ai-note/app/ai-note/internal/cmd"
	_ "github.com/wangle201210/ai-note/app/ai-note/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}

//
// type Req struct {
// 	Amount float64 `json:"amount"`
// 	Type   string  `json:"type"`
// 	Org    string  `json:"org"`
// }
//
// func main() {
// 	s := g.Server()
// 	s.BindHandler("/", func(r *ghttp.Request) {
// 		req := &Req{}
// 		r.Parse(req)
// 		data := fmt.Sprintf("假装我在写入数据库：%+v", req)
// 		fmt.Printf("假装我正在写数据到mysql: %+v\n", req)
// 		r.Response.Write(data)
// 	})
// 	s.SetPort(8000)
// 	s.Run()
// }
