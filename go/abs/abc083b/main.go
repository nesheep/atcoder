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

	var n, a, b int
	fmt.Fscan(r, &n, &a, &b)

	var ans int
	for i := 1; i <= n; i++ {
		var sum int
		tmp := i
		for tmp != 0 {
			sum += tmp % 10
			tmp /= 10
		}
		if sum >= a && sum <= b {
			ans += i
		}
	}

	fmt.Fprintln(w, ans)
}
