package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, y int
	fmt.Fscan(rd, &n, &y)

	res10000, res5000, res1000 := -1, -1, -1

Loop:
	for a := 0; a <= n; a++ {
		for b := 0; b <= n-a; b++ {
			c := n - a - b
			total := 10000*a + 5000*b + 1000*c
			if total == y {
				res10000 = a
				res5000 = b
				res1000 = c
				break Loop
			}
		}
	}

	fmt.Fprintln(wr, res10000, res5000, res1000)
}
