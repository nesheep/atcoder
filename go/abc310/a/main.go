package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, p, q int
	fmt.Fscan(rd, &n, &p, &q)

	var d int
	dMin := 1 << 60
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d)
		dMin = min(dMin, d)
	}

	fmt.Println(min(p, q+dMin))
}
