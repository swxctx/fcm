package main

import (
	"github.com/swxctx/fcm"
)

func main() {
	apiKey := ""
	deviceToken := ""

	// new client
	fcm.New(apiKey)

	fcm.Send(&fcm.PushMsg{
		To: deviceToken,
		Notification: fcm.NotificationPayload{
			Title: "通知测试",
			Body:  "你有一条新的消息",
		},
	})
}
