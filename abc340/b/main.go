package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var q int
	fmt.Fscan(rd, &q)

	a := make([]int, 0, q)
	for i := 0; i < q; i++ {
		var t, x int
		fmt.Fscan(rd, &t, &x)
		if t == 1 {
			a = append(a, x)
		} else {
			fmt.Println(a[len(a)-x])
		}
	}
}
