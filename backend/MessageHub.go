package main

import (
	"chatroom-backend/def"
	"chatroom-backend/ringqueue"
	"sync"
)

// 线程安全的消息广播、存储工具类。
type messageHub struct {
	// 大锁，每次操作前都要获取这个锁
	lock sync.Mutex

	// 历史消息
	msgRec *ringqueue.RingQueue[def.Message]

	// 每次 notifySet 更新时，此字段应 +1
	notifySetVersion int64

	// 当有人发送新消息时，在这里通知其它所有 goroutine
	notifySet map[chan *def.Message]struct{}

	// 当前用户信息列表，用户退出后不删除
	users map[def.UserId]def.User

	// 最小的未使用的用户 id（用户 id 从 0 开始计数，故初始值为 0）
	minUnusedId def.UserId
}

func NewMessageHub() *messageHub {
	return &messageHub{
		msgRec:    ringqueue.NewRingQueue[def.Message](def.MAX_MESSAGE_HISTORY),
		notifySet: make(map[chan *def.Message]struct{}),
		users:     make(map[def.UserId]def.User),
	}
}

func (h *messageHub) GetHistoryMessages(since int) (msgs []*def.Message, tailIndex int) {
	h.lock.Lock()
	defer h.lock.Unlock()

	return h.msgRec.GetSince(since)
}

func (h *messageHub) SubmitAndBroadcastMessage(msg *def.Message) {
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
