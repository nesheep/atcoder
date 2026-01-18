package main

import (
	"container/heap"
	"fmt"
)

func main() {
	_, _, g, w := input()
	dist := dijkstra(g, w)
	output(dist)
}

func dijkstra(g [][]int, w map[[2]int]int) []int {
	dist := make([]int, len(g))
	for i := range dist {
		dist[i] = -1
	}
	dist[0] = 0

	fixed := make(map[int]bool, len(g))

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{value: 0, priority: 1})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		v := item.value

		if fixed[v] {
			continue
		}

		fixed[v] = true

		for _, u := range g[v] {
			if fixed[u] {
				continue
			}
			d := dist[v] + w[[2]int{v, u}]
			if dist[u] == -1 || d < dist[u] {
				dist[u] = d
				heap.Push(&pq, &Item{value: u, priority: d})
			}
		}
	}

	return dist
}

func input() (int, int, [][]int, map[[2]int]int) {
	var n, m int
	fmt.Scan(&n, &m)
	g := make([][]int, n)
	w := make(map[[2]int]int, m)
	for range m {
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		a, b = a-1, b-1
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
		w[[2]int{a, b}] = c
		w[[2]int{b, a}] = c
	}
	return n, m, g, w
}

func output(dist []int) {
	for _, v := range dist {
		fmt.Println(v)
	}
}

type Item struct {
	value    int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
