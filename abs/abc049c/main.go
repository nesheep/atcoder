package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var s string
	fmt.Fscan(rd, &s)

	ds := []string{"dream", "dreamer", "erase", "eraser"}

	ans := true
	for i := len(s); i > 0; {
		var f bool
		for _, d := range ds {
			if strings.HasSuffix(s[:i], d) {
				f = true
				i -= len(d)
				break
			}
		}
		if !f {
			ans = false
			break
		}
	}

	if ans {
		fmt.Fprintln(wr, "YES")
	} else {
		fmt.Fprintln(wr, "NO")
	}
}
