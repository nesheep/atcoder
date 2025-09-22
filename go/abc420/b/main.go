package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n, m, s := input()
	p := calcPoints(n, m, s)
	ans := maxIndex(p)
	output(ans)
}

func input() (int, int, [][]byte) {
	var n, m int
	fmt.Scan(&n, &m)

	s := make([][]byte, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	return n, m, s
}

func calcPoints(n, m int, s [][]byte) []int {
	p := make([]int, n)
	for j := 0; j < m; j++ {
		var z, o int
		for i := 0; i < n; i++ {
			switch s[i][j] {
			case '0':
				z++
			case '1':
				o++
			}
		}
		for i := 0; i < n; i++ {
			if z < o && s[i][j] == '0' {
				p[i]++
			}
			if z > o && s[i][j] == '1' {
				p[i]++
			}
		}
	}
	return p
}

func maxIndex(a []int) []int {
	var v []int
	m := max(a)
	for i, x := range a {
		if x == m {
			v = append(v, i)
		}
	}
	return v
}

func max(a []int) int {
	var v int
	for _, x := range a {
		if x > v {
			v = x
		}
	}
	return v
}

func output(ans []int) {
	s := make([]string, 0, len(ans))
	for _, a := range ans {
		s = append(s, strconv.Itoa(a+1))
	}
	fmt.Println(strings.Join(s, " "))
}
