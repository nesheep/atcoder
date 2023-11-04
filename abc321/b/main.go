package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, x int
	fmt.Fscan(rd, &n, &x)

	min := 100
	max := 0
	sum := 0
	for i := 0; i < n-1; i++ {
		var a int
		fmt.Fscan(rd, &a)
		if a < min {
			min = a
		}
		if a > max {
			max = a
		}
		sum += a
	}

	if sum-min < x {
		fmt.Println(-1)
	} else if sum-max >= x {
		fmt.Println(0)
	} else {
		fmt.Println(x - (sum - min - max))
	}
}
