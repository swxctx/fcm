# fcm
Fcm for Golang

## Usage

```shell
go get github.com/swxctx/fcm
```

## Doc
#### Firebase Cloud Messaging HTTP Protocol Specs

```js
https://firebase.google.com/docs/cloud-messaging/http-server-ref
```

#### Firebase Cloud Messaging Developer docs

```js
https://firebase.google.com/docs/cloud-messaging/
```
#### (Google) Instance Id Server Reference

```js
https://developers.google.com/instance-id/reference/server
```

## Example

```go
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
```
