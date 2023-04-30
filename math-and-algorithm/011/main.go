package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	ps := make([]int, 0, n)
	for i := 2; i <= n; i++ {
		isPN := true
		for _, p := range ps {
			if i%p == 0 {
				isPN = false
				break
			}
		}
		if isPN {
			ps = append(ps, i)
		}
	}

	psa := make([]string, 0, len(ps))
	for _, p := range ps {
		psa = append(psa, strconv.Itoa(p))
	}

	fmt.Println(strings.Join(psa, " "))
}
