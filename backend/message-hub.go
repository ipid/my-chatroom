package main

import (
	"sync"
	"time"
)

// 线程安全的消息广播、存储工具类。
type messageHub struct {
	msgRecLock sync.Mutex
	msgRec     *MessageRecorder

	subscriberLock sync.RWMutex
	subscribers    map[chan *Message]struct{}
}

func NewMessageHub() *messageHub {
	return &messageHub{
		msgRecLock: sync.Mutex{},
		msgRec:     NewMessageRecorder(MAX_MESSAGE_HISTORY),

		subscriberLock: sync.RWMutex{},
		subscribers:    make(map[chan *Message]struct{}),
	}
}

func (h *messageHub) GetHistoryMessages(num int) []*Message {
	h.msgRecLock.Lock()
	defer h.msgRecLock.Unlock()

	return h.msgRec.GetLast(num)
}

func (h *messageHub) SendClientMessage(clientMsg *ClientMessage, senderChan chan *Message) {
	msg := &Message{
		Id:          -1,
		Sender:      clientMsg.Sender,
		Content:     clientMsg.Content,
		TimestampMs: uint64(time.Now().UnixMilli()),
	}

	h.msgRecLock.Lock()
	msg.Id = int64(h.msgRec.Put(msg))
	h.msgRecLock.Unlock()

	h.subscriberLock.RLock()
	defer h.subscriberLock.RUnlock()

	for subscriber := range h.subscribers {
		if subscriber == senderChan {
			continue
		}

		subscriber <- msg
	}
}

func (h *messageHub) Subscribe() chan *Message {
	h.subscriberLock.Lock()
	defer h.subscriberLock.Unlock()

	c := make(chan *Message, MESSAGE_BUFFER)
	h.subscribers[c] = struct{}{}

	return c
}

func (h *messageHub) Unsubscribe(msgChan chan *Message) {
	h.subscriberLock.Lock()
	defer h.subscriberLock.Unlock()

	delete(h.subscribers, msgChan)
}
