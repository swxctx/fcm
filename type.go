package fcm

/**
doc: https://firebase.google.com/docs/cloud-messaging/http-server-ref?hl=zh-cn#send-downstream
*/

// PushMsg
type PushMsg struct {
	// 值可以是设备的注册令牌、设备组的通知键或单个主题（前缀为 /topics/）。如需发送到多个主题，请使用 condition 参数。
	To string `json:"to,omitempty"`
	// 值应该是向其发送多播消息的目标注册令牌数组。该数组必须包含 1 到 1000 个注册令牌。如需向单一设备发送消息，请使用 to 参数。
	RegistrationIds []string `json:"registration_ids,omitempty"`
	// 参数指定用于确定消息目标的逻辑条件表达式。
	// 支持的条件：主题，格式为“'yourTopic' in topics”。此值不区分大小写。
	// 支持的运算符：&&、||。每个主题消息最多支持两个运算符。
	Condition string `json:"condition,omitempty"`
	// 此参数用于指定一组可折叠的消息（例如含有 collapse_key: "Updates Available"），以便当恢复传送时只发送最后一条消息。
	// 这是为了避免当设备恢复在线状态或变为活跃状态时重复发送过多相同的消息。
	// 请注意，消息发送顺序并不固定。
	CollapseKey string `json:"collapse_key,omitempty"`
	// 此参数用于指定设备离线后消息在 FCM 存储空间中保留的时长（以秒为单位）。
	// 支持的最长存留时间为 4 周，默认值为 4 周。
	TimeToLive int `json:"time_to_live,omitempty"`
	// 此参数用于指定应用的软件包名称，其注册令牌必须匹配才能接收消息。
	RestrictedPackageName string `json:"restricted_package_name,omitempty"`
	// 此参数用于指定消息载荷的自定义键值对。
	// 例如，对于 data:{"score":"3x1"}:
	Data interface{} `json:"data,omitempty"`

	// 此参数用于指定用户可见的预定义通知载荷键值对。如需了解详情，请参阅“通知载荷支持”。
	Notification NotificationPayload `json:"notification,omitempty"`
}

// NotificationPayload
type NotificationPayload struct {
	// 通知标题
	Title string `json:"title,omitempty"`
	// 通知的正文文字
	Body string `json:"body,omitempty"`
	// 通知的渠道 ID（Android O 中的新功能）
	AndroidChannelID string `json:"android_channel_id,omitempty"`
	// 通知的图标。
	Icon string `json:"icon,omitempty"`
	// 设备收到通知时要播放的声音。
	// 支持 "default" 或应用中绑定的声音资源的文件名
	Sound string `json:"sound,omitempty"`
	// 主屏幕应用图标上的标志值。角标
	Badge string `json:"badge,omitempty"`
	// 用于替换抽屉式通知栏中现有通知的标识符。
	Tag string `json:"tag,omitempty"`
	// 通知的图标颜色，以 #rrggbb 格式来表示
	Color string `json:"color,omitempty"`
	// 与用户点击通知相关的操作。
	// 如果已指定此参数，则当用户点击通知时，系统将会启动带有匹配 intent 过滤器的 Activity。
	ClickAction string `json:"click_action,omitempty"`
	// 应用的字符串资源中正文字符串的键，用于将正文文字本地化为用户当前的本地化设置语言。
	BodyLocKey string `json:"body_loc_key,omitempty"`
	// 将用于替换 body_loc_key（用来将正文文字本地化为用户当前的本地化设置语言）中的格式说明符的变量字符串值。
	BodyLocArgs string `json:"body_loc_args,omitempty"`
	// 应用的字符串资源中标题字符串的键，用于将标题文字本地化为用户当前的本地化设置语言。
	TitleLocKey string `json:"title_loc_key,omitempty"`
	// 将用于替换 title_loc_key（用来将标题文字本地化为用户当前的本地化设置语言）中的格式说明符的变量字符串值。
	TitleLocArgs string `json:"title_loc_args,omitempty"`
}

// Response
type Response struct {
	Ok bool
	// 状态码
	StatusCode int
	// 用于识别多播消息的唯一 ID（数字）
	MulticastId int64 `json:"multicast_id"`
	// 成功的消息数
	Success int `json:"success"`
	Fail    int `json:"failure"`
	// 代表已处理消息状态的对象数组。对象的排列顺序与请求相同（即对于请求中的每个注册 ID，其结果的列出顺序与响应中的索引顺序相同）。
	// message_id：字符串，用于指定每条成功处理的消息的唯一 ID。
	// error：字符串，用于指定处理收件人的消息时发生的错误。
	Results []map[string]string `json:"results,omitempty"`
	// 消息ID
	MsgId int64 `json:"message_id,omitempty"`
	// 错误信息
	Err string `json:"error,omitempty"`
}
