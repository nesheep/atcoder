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

	var a, x, ans uint64
	x = 1
	for i := 0; i < 64; i++ {
		fmt.Fscan(rd, &a)
		if a == 1 {
			ans += x
		}
		x *= 2
	}

	fmt.Fprintln(wr, ans)
}
