package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var x1, y1, r1 float64
	var x2, y2, r2 float64
	fmt.Fscan(r, &x1, &y1, &r1, &x2, &y2, &r2)

	a := math.Abs(r1 - r2)
	b := r1 + r2
	d := math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))

	var ans int
	switch {
	case d < a:
		ans = 1
	case d == a:
		ans = 2
	case a < d && d < b:
		ans = 3
	case d == b:
		ans = 4
	case b < d:
		ans = 5
	}

	fmt.Println(ans)
}
