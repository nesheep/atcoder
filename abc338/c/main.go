package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(rd, &n)

	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &q[i])
	}

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &a[i])
	}

	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &b[i])
	}

	aMax := 1 << 60
	for i := 0; i < n; i++ {
		if a[i] == 0 {
			continue
		}
		p := q[i] / a[i]
		if p < aMax {
			aMax = p
		}
	}

	var ans int
	for i := 0; i <= aMax; i++ {
		r := make([]int, n)
		for j := 0; j < n; j++ {
			r[j] = q[j] - a[j]*i
		}
		bMax := 1 << 60
		for j := 0; j < n; j++ {
			if b[j] == 0 {
				continue
			}
			p := r[j] / b[j]
			if p < bMax {
				bMax = p
			}
		}
		if i+bMax > ans {
			ans = i + bMax
		}
	}

	fmt.Println(ans)
}
