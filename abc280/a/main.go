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

	var ans int
	for i := 0; i < h; i++ {
		var s []byte
		fmt.Fscan(rd, &s)
		for _, v := range s {
			if v == '#' {
				ans++
			}
		}
	}

	fmt.Fprintln(wr, ans)
}
