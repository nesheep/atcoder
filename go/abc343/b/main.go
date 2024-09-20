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

	var n int
	fmt.Fscan(rd, &n)

	for i := 0; i < n; i++ {
		a := make([]any, 0, n)
		for j := 0; j < n; j++ {
			var b int
			fmt.Fscan(rd, &b)
			if b != 0 {
				a = append(a, j+1)
			}
		}
		fmt.Fprintln(wr, a...)
	}
}
