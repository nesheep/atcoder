package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(r, &n)

	xs := make([]float64, n*5)
	for i := 0; i < n*5; i++ {
		fmt.Fscan(r, &xs[i])
	}

	sort.Float64s(xs)

	var sum float64
	for _, v := range xs[n : len(xs)-n] {
		sum += v
	}

	fmt.Printf("%.15f\n", sum/float64(n*3))
}
