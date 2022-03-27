package main

// MessageRecorder 基于环形队列实现一个历史消息记录功能。
// 注意：本类并不是线程安全的，请自行加锁。
type MessageRecorder struct {
	histories       []*Message
	tail, tailIndex int
}

func NewMessageRecorder(length int) *MessageRecorder {
	return &MessageRecorder{
		histories: make([]*Message, length),
		tail:      length - 1,
		tailIndex: -1,
	}
}

func (cb *MessageRecorder) GetLast(maxMessageNum int) (lastMessages []*Message) {
	l := len(cb.histories)

	actualMessageNum := maxMessageNum
	if actualMessageNum > cb.tailIndex+1 {
		actualMessageNum = cb.tailIndex + 1
	}
	if actualMessageNum > l {
		actualMessageNum = l
	}

	lastMessages = make([]*Message, actualMessageNum)
	i, historyIndex := actualMessageNum-1, cb.tail+l
	for i >= 0 {
		lastMessages[i] = cb.histories[historyIndex%l]
		i--
		historyIndex--
	}

	return
}

func (cb *MessageRecorder) Put(msg *Message) (msgIndex int) {
	cb.tail = (cb.tail + 1) % len(cb.histories)
	cb.tailIndex++
	cb.histories[cb.tail] = msg
	return cb.tailIndex
}
