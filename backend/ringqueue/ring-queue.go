package ringqueue

// RingQueue 环形队列。
// 注意：本类并不是线程安全的，请自行加锁。
type RingQueue[T any] struct {
	histories       []*T
	tail, tailIndex int
}

// NewRingQueue 创建一个新的环形队列，其容量为 length。
func NewRingQueue[T any](length int) *RingQueue[T] {
	return &RingQueue[T]{
		histories: make([]*T, length),
		tail:      length - 1,
		tailIndex: -1,
	}
}

// GetSince 获取队列中所有虚拟下标大于等于 index 的元素。
// 当 index 小于最小元素下标时，返回结果中不会有超出范围的元素；index 太大时会返回空切片。
func (cb *RingQueue[T]) GetSince(index int) (elements []*T, tailIndex int) {
	// 环形队列的容量是一个重要的信息
	l := len(cb.histories)

	// 计算实际上需要返回多少个元素
	actualSize := cb.tailIndex - index
	if actualSize <= 0 {
		return nil, 0
	}
	if actualSize > l {
		actualSize = l
	}

	// 将 RingQueue 中的元素逐个拷贝到 elements 切片上
	elements = make([]*T, actualSize)
	i, historyIndex := actualSize-1, cb.tail+l
	for i >= 0 {
		elements[i] = cb.histories[historyIndex%l]
		i--
		historyIndex--
	}

	tailIndex = cb.tailIndex
	return
}

// Put 往环形队列中添加元素，并返回该元素的虚拟下标。
func (cb *RingQueue[T]) Put(msg *T) (msgIndex int) {
	cb.tail = (cb.tail + 1) % len(cb.histories)
	cb.tailIndex++
	cb.histories[cb.tail] = msg
	return cb.tailIndex
}
