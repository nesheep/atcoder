package main

import (
	"bufio"
	"fmt"
	"os"
)

type Queue struct {
	data [][2]int
	size int
}

func NewQueue(cap int) *Queue {
	return &Queue{data: make([][2]int, 0, cap), size: 0}
}

func (q *Queue) Push(n [2]int) {
	q.data = append(q.data, n)
	q.size++
}

func (q *Queue) Pop() bool {
	if q.IsEmpty() {
		return false
	}
	q.size--
	q.data = q.data[1:]
	return true
}

func (q *Queue) Front() [2]int {
	return q.data[0]
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func main() {
	rd := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscan(rd, &h, &w)

	s := make([][]byte, h)
	for i := range s {
		fmt.Fscan(rd, &s[i])
	}

	m := make([][]bool, h)
	for i := range m {
		m[i] = make([]bool, w)
	}

	var ans int

	q := NewQueue(h * w)
	for i := range s {
		for j := range s[i] {
			if m[i][j] {
				continue
			}
			m[i][j] = true
			if s[i][j] != '#' {
				continue
			}
			ans++
			q.Push([2]int{i, j})
			for !q.IsEmpty() {
				p := q.Front()
				q.Pop()
				for k := -1; k <= 1; k++ {
					for l := -1; l <= 1; l++ {
						if k == 0 && l == 0 {
							continue
						}
						x, y := p[0]+k, p[1]+l
						if x < 0 || x >= h || y < 0 || y >= w {
							continue
						}
						if m[x][y] {
							continue
						}
						m[x][y] = true
						if s[x][y] == '#' {
							q.Push([2]int{x, y})
						}
					}
				}
			}
		}
	}

	fmt.Println(ans)
}
