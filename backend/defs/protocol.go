package defs

type LoginToken string

// WebSocketParam 连接后端 WebSocket 服务器时，在 URL 的 query string 处添加上的内容
type WebSocketParam struct {
	// 上次登录用的用户 id
	UserId UserId `json:"userId"`

	// 上次登录时服务器发放的 token
	Token LoginToken `json:"token"`
}

// Request 客户端请求，具体类型由哪个字段存在而决定
type Request struct {
	// 客户端请求，若此字段存在，表示客户端希望发送一条消息
	Send *struct {
		// 要发送的消息内容，可以是普通文本，也可以是 1d9 等骰子文本
		Content string `json:"content"`
	} `json:"send,omitempty"`

	// 客户端请求，若此字段存在，表示客户端希望修改当前用户的显示名称
	Rename *struct {
		// 新的显示名称
		NewName string `json:"newName"`
	} `json:"rename,omitempty"`
}

// Response 表示一条服务器端的消息，具体类型由哪个字段存在而决定
type Response struct {
	// 服务器端消息回应，表示收到了一或多条消息
	Msgs []Message `json:"msgs,omitempty"`

	// 服务器端确认回应，表示客户端发出的请求得到了服务器端的确认
	Ack *Message `json:"ack,omitempty"`

	// 服务器端初始回应，用于将 token、当前用户、当前历史消息等告诉用户
	Init *struct {
		// 客户端当前的用户 id
		UserId UserId `json:"userId"`

		// 第一次登录时，这里会存放登录用的 token；用 token 登录时此字段为空字符串
		Token LoginToken `json:"token"`

		// 服务器端的历史消息
		History []Message `json:"history"`
	} `json:"init,omitempty"`

	Rename *struct {
		// 改名操作是否成功
		Success bool `json:"success"`

		// 若改名失败，则此字段表明失败原因
		Reason *RenameFailReason `json:"reason,omitempty"`
	} `json:"rename,omitempty"`
}

// RenameFailReason 改名操作失败的原因，该原因由服务器端产生
type RenameFailReason string

const (
	// CONFLICT 名字与他人重复，造成冲突
	CONFLICT RenameFailReason = "conflict"
)
