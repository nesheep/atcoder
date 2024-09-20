package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, q int
	fmt.Fscan(rd, &n, &q)

	ts := make([]bool, n)
	for i := range ts {
		ts[i] = true
	}

	for i := 0; i < q; i++ {
		var t int
		fmt.Fscan(rd, &t)
		t--
		ts[t] = !ts[t]
	}

	var ans int
	for i := range ts {
		if ts[i] {
			ans++
		}
	}

	fmt.Println(ans)
}
