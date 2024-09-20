package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b int) int {
	var f func(a, b int) int
	f = func(a, b int) int {
		if b == 0 {
			return a
		}
		return f(b, a%b)
	}
	if a >= b {
		return f(a, b)
	}
	return f(b, a)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	ans := a[0]
	for i := 1; i < n; i++ {
		ans = gcd(ans, a[i])
	}

	fmt.Fprintln(w, ans)
}
