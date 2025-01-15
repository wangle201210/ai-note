package dify

import (
	"context"
	"github.com/KevinZhao/dify-sdk-go"
	"github.com/gogf/gf/v2/frame/g"
	"sync"
)

var conversationMap = &sync.Map{}

var difyClient = dify.NewClient("http://localhost", "app-TqDfKPBA7mUiUeTTdf1C7MMv")

// var difyClient = dify.NewClientWithConfig(&dify.ClientConfig{
// 	Host:             "https://api.dify.ai",
// 	DefaultAPISecret: "app-he2abcBb4KJBaQxcX9slGM45",
// 	Timeout:          time.Second * 15,
// })

func Dify(req *dify.ChatMessageRequest) (err error, resp *dify.ChatMessageResponse) {
	value, ok := conversationMap.Load(req.User)
	if ok {
		req.ConversationID = value.(string)
	}
	resp, err = difyClient.Api().ChatMessages(context.Background(), req)
	if err != nil {
		g.Log().Errorf(context.Background(), "Dify err %+v", err)
		return
	}
	g.Log().Infof(context.Background(), "Dify data %+v", resp)
	conversationMap.Store(req.User, resp.ConversationID)
	return
}
