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
	var s []byte
	fmt.Fscan(rd, &n, &s)

	for i := 0; i < n-1; i++ {
		l := n - i - 1
		for k := 0; k < n-i-1; k++ {
			if s[k] == s[k+i+1] {
				l = k
				break
			}
		}
		fmt.Fprintln(wr, l)
	}
}
