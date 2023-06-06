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

	var n, m int
	fmt.Fscan(rd, &n, &m)

	ss := make([][]byte, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &ss[i])
	}

	var ans int
	for i, sa := range ss {
		for _, sb := range ss[i+1:] {
			f := true
			for j := 0; j < m; j++ {
				if sa[j] != 'o' && sb[j] != 'o' {
					f = false
					break
				}
			}
			if f {
				ans++
			}
		}
	}

	fmt.Fprintln(wr, ans)
}
