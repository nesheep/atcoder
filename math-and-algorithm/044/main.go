package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, m int
	fmt.Fscan(rd, &n, &m)

	g := make([][]int, n+1)
	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscan(rd, &a, &b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	dist := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dist[i] = -1
	}

	q := NewQueueInt(0)
	q.Push(1)
	for !q.IsEmpty() {
		pos := q.Front()
		q.Pop()
		for _, nex := range g[pos] {
			if dist[nex] == -1 {
				dist[nex] = dist[pos] + 1
				q.Push(nex)
			}
		}
	}

	for _, d := range dist[1:] {
		fmt.Fprintln(wr, d)
	}
}
