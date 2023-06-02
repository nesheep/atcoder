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

	ss := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &ss[i])
	}

	for i := n - 1; i >= 0; i-- {
		fmt.Fprintln(wr, ss[i])
	}
}
