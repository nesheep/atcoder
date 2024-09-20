package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, t int
	fmt.Fscan(rd, &n, &t)

	mr := make(map[int]int, n)
	mc := make(map[int]int, n)
	md := make(map[int]int, 2)

	for i := 0; i < t; i++ {
		var a int
		fmt.Fscan(rd, &a)

		r := (a - 1) / n
		mr[r]++

		c := (a - 1) % n
		mc[c]++

		if r == c {
			md[0]++
		}

		if r == n-c-1 {
			md[1]++
		}

		if mr[r] == n || mc[c] == n || md[0] == n || md[1] == n {
			fmt.Println(i + 1)
			os.Exit(0)
		}
	}

	fmt.Println(-1)
}
