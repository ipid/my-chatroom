package main

import "unsafe"

const MAX_MESSAGE_HISTORY = 10
const MESSAGE_BUFFER = 16

type Message struct {
	Id          int64  `json:"id"`
	Sender      string `json:"sender"`
	Content     string `json:"content"`
	TimestampMs uint64 `json:"timestampMs"`
}

type ClientMessage struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
}

var a unsafe.Pointer