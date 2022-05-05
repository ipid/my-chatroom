package main

import (
	"encoding/json"
)

const MAX_MESSAGE_HISTORY = 10

/* 客户端请求类型 */
const REQ_TYPE_MESSAGE = "send-msg"

/* 服务端返回类型 */
const RESP_TYPE_NEW_MESSAGE = "msg"
const RESP_TYPE_MESSAGE_ACKNOWLEDGE = "msg-ack"

type UnknownClientRequest struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type ClientRequest struct {
	Type       string
	MsgRequest *ClientRequestMessage
}

type ClientRequestMessage struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
}

type Message struct {
	Id          int64  `json:"id"`
	Sender      string `json:"sender"`
	SenderId    int64  `json:"-"`
	TimestampMs int64  `json:"timestampMs"`

	// 当此字段不为 nil 时，表明该消息为富文本类型消息
	RichText *MessageDataRichText `json:"richText,omitempty"`

	// 当此字段不为 nil 时，表明该消息为骰子类型消息
	Dice *MessageDataDice `json:"dice,omitempty"`
}

type MessageDataRichText struct {
	Content string `json:"content"`
}

type MessageDataDice struct {
	Dices []int `json:"dices"`
}

type ServerResponse struct {
	Type string   `json:"type"`
	Data *Message `json:"data"`
}
