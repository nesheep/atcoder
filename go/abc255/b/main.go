package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const INF = 1 << 60

type Point struct {
	X, Y int
}

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(rd, &n, &k)

	a := make(map[int]bool, k)
	for i := 0; i < k; i++ {
		var ai int
		fmt.Fscan(rd, &ai)
		a[ai] = true
	}

	p := make(map[int]Point, n)
	for i := 1; i <= n; i++ {
		var x, y int
		fmt.Fscan(rd, &x, &y)
		p[i] = Point{x, y}
	}

	var ans int
	for i := range p {
		if a[i] {
			continue
		}
		dMin := INF
		for j := range a {
			dx := p[i].X - p[j].X
			dy := p[i].Y - p[j].Y
			d := dx*dx + dy*dy
			if d < dMin {
				dMin = d
			}
		}
		if dMin > ans {
			ans = dMin
		}
	}

	fmt.Printf("%.12f\n", math.Sqrt(float64(ans)))
}
