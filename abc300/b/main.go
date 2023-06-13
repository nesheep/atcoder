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

	var h, w int
	fmt.Fscan(rd, &h, &w)

	a := make([][]byte, h)
	for i := range a {
		fmt.Fscan(rd, &a[i])
	}

	b := make([][]byte, h)
	for i := range b {
		fmt.Fscan(rd, &b[i])
	}

	var ok bool
All:
	for s := 0; s < h; s++ {
		for t := 0; t < w; t++ {
			ok = true
		One:
			for i := 0; i < h; i++ {
				for j := 0; j < w; j++ {
					if a[(i+s)%h][(j+t)%w] != b[i][j] {
						ok = false
						break One
					}
				}
			}
			if ok {
				break All
			}
		}
	}

	if ok {
		fmt.Fprintln(wr, "Yes")
	} else {
		fmt.Fprintln(wr, "No")
	}
}
