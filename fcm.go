package fcm

var (
	client *Client

	// 接口地址
	apiServerUrl = "https://fcm.googleapis.com/fcm/send"
)

// Client
type Client struct {
	ApiKey string
}

/*
  New
  @Desc: new fcm push client Cloud Messaging API 服务器密钥
  @param: apiKey
  @return: *Client
*/
func New(apiKey string) {
	client = &Client{
		ApiKey: apiKey,
	}
}

// Send 发送消息
func Send(pushMsg *PushMsg) (*Response, error) {
	return client.send(pushMsg)
}
