package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pair struct {
	X, Y int
}

type QueuePair struct {
	data []Pair
	size int
}

func NewQueuePair(cap int) *QueuePair {
	return &QueuePair{data: make([]Pair, 0, cap), size: 0}
}

func (q *QueuePair) Push(n Pair) {
	q.data = append(q.data, n)
	q.size++
}

func (q *QueuePair) Pop() bool {
	if q.IsEmpty() {
		return false
	}
	q.size--
	q.data = q.data[1:]
	return true
}

func (q *QueuePair) Front() Pair {
	return q.data[0]
}

func (q *QueuePair) IsEmpty() bool {
	return q.size == 0
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var r, c, sy, sx, gy, gx int
	fmt.Fscan(rd, &r, &c, &sy, &sx, &gy, &gx)

	m := make([][]byte, r)
	for i := range m {
		fmt.Fscan(rd, &m[i])
	}

	p := make([][]int, r)
	for i := range p {
		p[i] = make([]int, c)
		for j := range p[i] {
			p[i][j] = -1
		}
	}

	q := NewQueuePair(r * c)
	p[sy-1][sx-1] = 0
	q.Push(Pair{sy - 1, sx - 1})
	for !q.IsEmpty() {
		v := q.Front()
		q.Pop()
		if v.X == gy-1 && v.Y == gx-1 {
			break
		}
		for _, d := range []Pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			if m[v.X+d.X][v.Y+d.Y] == '.' && p[v.X+d.X][v.Y+d.Y] == -1 {
				p[v.X+d.X][v.Y+d.Y] = p[v.X][v.Y] + 1
				q.Push(Pair{v.X + d.X, v.Y + d.Y})
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", p[gy-1][gx-1])
}
