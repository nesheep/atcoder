package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    int
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func main() {
	var k int
	fmt.Scan(&k)

	g := make([][]Item, k)
	for i := range g {
		for j := 0; j < 10; j++ {
			if i == 0 && j == 0 {
				continue
			}
			g[i] = append(g[i], Item{value: (i*10 + j) % k, priority: j})
		}
	}

	visited := make([]bool, k)
	dist := make([]int, k)
	for i := range dist {
		dist[i] = -1
	}

	pq := PriorityQueue{}
	heap.Init(&pq)
	heap.Push(&pq, &Item{value: 0, priority: 0})
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		if visited[item.value] {
			continue
		}
		visited[item.value] = true
		for _, next := range g[item.value] {
			to := next.value
			cost := next.priority
			if item.value != 0 {
				cost += dist[item.value]
			}
			if dist[to] == -1 || dist[to] > cost {
				dist[to] = cost
				heap.Push(&pq, &Item{value: to, priority: cost})
			}
		}
	}

	fmt.Println(dist[0])
}
