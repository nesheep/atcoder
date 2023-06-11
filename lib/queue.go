package lib

type QueueInt struct {
	data []int
	size int
}

func NewQueueInt(cap int) *QueueInt {
	return &QueueInt{data: make([]int, 0, cap), size: 0}
}

func (q *QueueInt) Push(n int) {
	q.data = append(q.data, n)
	q.size++
}

func (q *QueueInt) Pop() bool {
	if q.IsEmpty() {
		return false
	}
	q.size--
	q.data = q.data[1:]
	return true
}

func (q *QueueInt) Front() int {
	return q.data[0]
}

func (q *QueueInt) IsEmpty() bool {
	return q.size == 0
}
