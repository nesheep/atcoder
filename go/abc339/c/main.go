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

	a := make([]int, n+1)
	aMin := 0

	for i := 1; i <= n; i++ {
		var b int
		fmt.Fscan(rd, &b)
		a[i] = a[i-1] + b
		if a[i] < aMin {
			aMin = a[i]
		}
	}

	fmt.Println(a[len(a)-1] - aMin)
}
