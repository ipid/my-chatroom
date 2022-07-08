package def

// UserId 用户 ID（从 0 开始）
type UserId int64

// MsgId 消息 ID（从 0 开始）
type MsgId int64

// User 聊天室内的用户
type User struct {
	// 用户 id，在同一聊天室中唯一
	UserId UserId `json:"userId"`

	// 用户的显示名称
	Name string `json:"name"`
}

// Dice 骰子消息中的一个骰子
type Dice struct {
	// 表示该骰子的取值范围是 [1, Max]
	Max int `json:"max"`

	// 表示该骰子投出来的值
	Value int `json:"value"`
}

// Message 表示一条消息，该消息的类型由哪个字段存在而决定
type Message struct {
	// 消息 id
	MsgId MsgId `json:"msgId"`

	// 发送该消息的用户 id
	SenderUserId UserId `json:"senderUserId"`

	// 发送消息时的时间戳（单位：毫秒）
	TimestampMs int64 `json:"timestampMs"`

	// 文本消息
	Text *struct {
		// 表示文本消息的内容
		Content string `json:"content"`
	} `json:"text,omitempty"`

	// 骰子消息
	Dices []Dice `json:"dices,omitempty"`
}

// DiceRequest 表示投一到多个某种骰子的请求，如 1d99、3d6 等。
type DiceRequest struct {
	// 骰子个数
	Num int

	// 骰子的最大值。表示该骰子的取值范围是 [1, Max]
	Max int
}
