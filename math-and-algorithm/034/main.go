package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const inf = 1000000000000.0

func main() {
	r := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(r, &n)

	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &x[i])
		fmt.Fscan(r, &y[i])
	}

	ans := inf
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := math.Sqrt(math.Pow(x[j]-x[i], 2) + math.Pow(y[j]-y[i], 2))
			ans = math.Min(ans, d)
		}
	}

	fmt.Printf("%.12f\n", ans)
}
