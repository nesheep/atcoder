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

	s := make([][]byte, n)
	for i := range s {
		fmt.Fscan(rd, &s[i])
	}

	for i := 0; i < n; i++ {
	Loop:
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			t := make([]byte, 0, len(s[i])+len(s[j]))
			t = append(t, s[i]...)
			t = append(t, s[j]...)
			for k := 0; k < len(t)/2; k++ {
				if t[k] != t[len(t)-k-1] {
					continue Loop
				}
			}
			fmt.Fprintln(wr, "Yes")
			return
		}
	}

	fmt.Fprintln(wr, "No")
}
