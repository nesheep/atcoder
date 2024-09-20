package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	p := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		p[i] = true
	}

	for i := 2; i+i <= n; i++ {
		if p[i] {
			for j := i * 2; j <= n; j += i {
				p[j] = false
			}
		}
	}

	var m int
	for i := 2; i <= n; i++ {
		if p[i] {
			m++
		}
	}

	ans := make([]string, 0, m)
	for i := 2; i <= n; i++ {
		if p[i] {
			ans = append(ans, strconv.Itoa(i))
		}
	}

	fmt.Println(strings.Join(ans, " "))
}
