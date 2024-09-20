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

	a := make([][]byte, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &a[i])
	}

	ds := [][2]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}

	var m int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for d := range ds {
				sum := int(a[i][j] - '0')
				p := [2]int{i, j}
				for k := 0; k < n-1; k++ {
					p[0] += ds[d][0]
					p[1] += ds[d][1]
					if p[0] < 0 {
						p[0] = n - 1
					}
					if p[0] > n-1 {
						p[0] = 0
					}
					if p[1] < 0 {
						p[1] = n - 1
					}
					if p[1] > n-1 {
						p[1] = 0
					}
					sum *= 10
					sum += int(a[p[0]][p[1]] - '0')
				}
				if sum > m {
					m = sum
				}
			}
		}
	}

	fmt.Println(m)
}
