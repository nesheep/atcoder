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

	var n, k int
	fmt.Fscan(rd, &n, &k)

	b := make([]any, 0, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(rd, &a)
		if a%k == 0 {
			b = append(b, a/k)
		}
	}

	fmt.Fprintln(wr, b...)
}
