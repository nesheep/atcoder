package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n, d int
	fmt.Fscan(r, &n, &d)

	ts := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &ts[i])
	}

	ans := -1
	for i := 2; i <= n; i++ {
		if ts[i]-ts[i-1] <= d {
			ans = ts[i]
			break
		}
	}

	fmt.Println(ans)
}
