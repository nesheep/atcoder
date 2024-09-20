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

	var s, t []byte
	fmt.Fscan(rd, &s, &t)

	ans := len(t)
	for i := range s {
		if s[i] != t[i] {
			ans = i + 1
			break
		}
	}

	fmt.Fprintln(wr, ans)
}
