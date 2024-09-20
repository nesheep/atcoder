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

	var n, x int
	fmt.Fscan(rd, &n, &x)

	var p int
	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &p)
		if p == x {
			fmt.Fprintln(wr, i)
			return
		}
	}
}
