package dify

import (
	"github.com/KevinZhao/dify-sdk-go"
	"testing"
)

func TestDify(t *testing.T) {
	err, resp := Dify(&dify.ChatMessageRequest{
		User:  "wanna",
		Query: "荔枝8元一斤，我买了5斤，给了50，商家补了我3块",
	})
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	t.Logf("%+v", resp.Answer)
}
