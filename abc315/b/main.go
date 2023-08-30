package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var m int
	fmt.Fscan(rd, &m)

	sum := make([]int, m+1)
	for i := 1; i <= m; i++ {
		var d int
		fmt.Fscan(rd, &d)
		sum[i] = sum[i-1] + d
	}

	mid := sum[m]/2 + 1
	for i := 1; i <= m; i++ {
		if sum[i] >= mid {
			fmt.Println(i, mid-sum[i-1])
			os.Exit(0)
		}
	}
}
