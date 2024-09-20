package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(rd, &n, &m)

	g := make(map[[2]int]bool, m)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(rd, &u, &v)
		g[[2]int{u, v}] = true
		g[[2]int{v, u}] = true
	}

	var ans int
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			for k := j + 1; k <= n; k++ {
				if g[[2]int{i, j}] && g[[2]int{j, k}] && g[[2]int{k, i}] {
					ans++
				}
			}
		}
	}

	fmt.Println(ans)
}
