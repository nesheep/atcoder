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

	covered := make(map[[2]int]struct{})
	for i := 0; i < n; i++ {
		var a, b, c, d int
		fmt.Fscan(rd, &a, &b, &c, &d)
		for x := a; x < b; x++ {
			for y := c; y < d; y++ {
				covered[[2]int{x, y}] = struct{}{}
			}
		}
	}

	fmt.Println(len(covered))
}
