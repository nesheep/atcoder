package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n int
	fmt.Scan(&n)

	var h, maxH, ans int
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &h)
		if maxH <= h {
			maxH = h
			ans = i
		}
	}

	fmt.Println(ans + 1)
}
