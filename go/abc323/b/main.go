package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	players := make([]int, n)
	scores := make(map[int]int, n)

	s := make([][]byte, n)
	for i := range s {
		fmt.Scan(&s[i])
		players[i] = i + 1
		scores[players[i]] = bytes.Count(s[i], []byte{'o'})
	}

	sort.SliceStable(players, func(i, j int) bool {
		return scores[players[i]] > scores[players[j]]
	})

	ans := make([]string, n)
	for i := range players {
		ans[i] = strconv.Itoa(players[i])
	}

	fmt.Println(strings.Join(ans, " "))
}
