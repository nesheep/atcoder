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

	var a bool
	var ans uint64
	for i := 0; i < 64; i++ {
		fmt.Fscan(rd, &a)
		if a {
			ans += 1 << uint64(i)
		}
	}

	fmt.Fprintln(wr, ans)
}
