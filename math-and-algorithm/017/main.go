package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func GCD(a, b int) int {
	if a >= b {
		return gcd(a, b)
	}
	return gcd(b, a)
}

func LCM(a, b int) int {
	g := GCD(a, b)
	return a / g * b
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
		ans = LCM(ans, a[i])
	}

	fmt.Fprintln(w, ans)
}
