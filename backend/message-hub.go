package main

import (
	"github.com/ipid/chatroom-backend/ringqueue"
	"sync"
)

// 线程安全的消息广播、存储工具类。
type messageHub struct {
	msgRecLock sync.Mutex
	msgRec     *ringqueue.NewRingQueue

	subscriberLock sync.RWMutex
	lastUnusedId   int64
	subscribers    map[int64]chan *ServerResponse
}

func NewMessageHub() *messageHub {
	h := &messageHub{
		msgRecLock: sync.Mutex{},
		msgRec:     ringqueue.NewMessageRecorder(MAX_MESSAGE_HISTORY),

		subscriberLock: sync.RWMutex{},
		lastUnusedId:   0,
		subscribers:    make(map[int64]chan *ServerResponse),
	}
	return h
}

func (h *messageHub) GetHistoryMessages(num int) []*Message {
	h.msgRecLock.Lock()
	defer h.msgRecLock.Unlock()

	return h.msgRec.GetLast(num)
}

func (h *messageHub) SubmitAndBroadcastMessage(msg *Message) {
	h.msgRecLock.Lock()
	msg.Id = int64(h.msgRec.Put(msg))
	h.msgRecLock.Unlock()

	h.subscriberLock.RLock()
	defer h.subscriberLock.RUnlock()

	resp := &ServerResponse{
		Type: RESP_TYPE_NEW_MESSAGE,
		Data: msg,
	}

	for id, subscriber := range h.subscribers {
		if id == msg.SenderId {
			continue
		}

		subscriber <- resp
	}
}

func (h *messageHub) Subscribe() (chanWriteToClient <-chan *ServerResponse, id int64) {
	h.subscriberLock.Lock()
	defer h.subscriberLock.Unlock()

	id = h.lastUnusedId
	h.lastUnusedId++

	h.subscribers[id] = make(chan *ServerResponse)
	return h.subscribers[id], id
}

func (h *messageHub) Unsubscribe(id int64) {
	h.subscriberLock.Lock()
	defer h.subscriberLock.Unlock()

	if c, has := h.subscribers[id]; has {
		close(c)
		delete(h.subscribers, id)
	}
}
