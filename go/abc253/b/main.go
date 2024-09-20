package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	rd := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscan(rd, &h, &w)

	k := make([][2]int, 0, 2)
	for i := 0; i < h; i++ {
		var s []byte
		fmt.Fscan(rd, &s)
		for j := 0; j < w; j++ {
			if s[j] == 'o' {
				k = append(k, [2]int{i, j})
			}
		}
	}

	fmt.Println(abs(k[0][0]-k[1][0]) + abs(k[0][1]-k[1][1]))
}
