package wechat

import (
	"fmt"
	"github.com/KevinZhao/dify-sdk-go"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	dc "github.com/wangle201210/ai-note/app/ai-note/internal/controller/dify"
	"time"
)

var oa *officialaccount.OfficialAccount
var mem = cache.NewMemory()

func init() {
	cfg := &offConfig.Config{
		AppID:          "wx47467701add606b4",
		AppSecret:      "890bd56666b6e661c0398269ed08e527",
		Token:          "noteai",
		EncodingAESKey: "p5vHZ32j1cS2L3404I5tI7dSGbhe8rhwBPqXswPwjxr",
		Cache:          mem,
	}
	wc := wechat.NewWechat()
	oa = wc.GetOfficialAccount(cfg)
}

func getMKey(msgID int64) string {
	return fmt.Sprintf("msg:%d", msgID)
}

func getMResKey(msgID int64) string {
	return fmt.Sprintf("res:%d", msgID)
}

func Server(r *ghttp.Request) {
	req := r.Request
	rw := r.Response.Writer

	// 传入request和responseWriter
	server := oa.GetServer(req, rw)
	// 设置接收消息的处理方法
	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		if mem.IsExist(getMKey(msg.MsgID)) {
			// 有过回复就直接回复
			if mem.IsExist(getMResKey(msg.MsgID)) {
				res := mem.Get(getMResKey(msg.MsgID)).(string)
				defer func() {
					mem.Delete(getMResKey(msg.MsgID))
				}()
				g.Log().Debugf(r.Context(), "msg=%+v,已经处理过了,并且有结果：%+v", msg.MsgID, res)
				return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText(res)}
			}
			g.Log().Debugf(r.Context(), "msg=%+v,已经处理过了", msg.MsgID)
			return nil
		}
		err := mem.Set(getMKey(msg.MsgID), 1, time.Minute)
		if err != nil {
			g.Log().Errorf(r.Context(), "Set Mem err %+v", err)
			return nil
		}
		g.Log().Debugf(r.Context(), "msg=%+v", msg)
		if msg.Event == message.EventSubscribe {
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText(
				"我是记账小助手，请输入你的消费情况，我将自动进行分类，输入示例1: 包子3.2，薄饼2.5，鸡肉1.8、输入示例2: 今日消费情况")}
		}
		// 回复消息：演示回复用户发送的消息
		reply := ""
		err, resp := dc.Dify(&dify.ChatMessageRequest{
			User:  msg.GetOpenID(),
			Query: msg.Content,
		})
		if err != nil {
			reply = err.Error()
		} else {
			reply = resp.Answer
			mem.Set(getMResKey(msg.MsgID), reply, time.Minute)
		}
		// msg.OpenID
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText(reply)}
	})

	// 处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 发送回复的消息
	server.Send()
}
