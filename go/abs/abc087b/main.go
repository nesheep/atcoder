package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var a, b, c, x int
	fmt.Fscan(r, &a, &b, &c, &x)

	var ans int
	for ai := 0; ai <= a; ai++ {
		for bi := 0; bi <= b; bi++ {
			l := x - 500*ai - 100*bi
			if l >= 0 && l <= 50*c {
				ans++
			}
		}
	}

	fmt.Fprintln(w, ans)
}
