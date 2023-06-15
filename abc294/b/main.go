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

	var a int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fscan(rd, &a)
			if a == 0 {
				fmt.Fprint(wr, ".")
			} else {
				fmt.Fprintf(wr, "%c", 'A'+a-1)
			}
		}
		fmt.Fprintln(wr)
	}
}
