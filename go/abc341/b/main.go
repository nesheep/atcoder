package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(rd, &n)

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(rd, &a[i])
	}

	var s, t int
	for i := 0; i < n-1; i++ {
		fmt.Fscan(rd, &s, &t)
		a[i+1] += t * (a[i] / s)
	}

	fmt.Println(a[n-1])
}
