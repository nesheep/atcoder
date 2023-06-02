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

	var t int
	fmt.Fscan(rd, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(rd, &n)

		var a, ans int
		for j := 0; j < n; j++ {
			fmt.Fscan(rd, &a)
			if a%2 == 1 {
				ans++
			}
		}

		fmt.Fprintln(wr, ans)
	}
}
