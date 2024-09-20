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

	nMax := 100
	a := make([]int, 0, nMax)
	for i := 0; i < nMax; i++ {
		var ai int
		fmt.Fscan(rd, &ai)
		a = append(a, ai)
		if ai == 0 {
			break
		}
	}

	for i := 0; i < len(a); i++ {
		fmt.Fprintln(wr, a[len(a)-i-1])
	}
}
