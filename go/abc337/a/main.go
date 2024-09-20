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

	var sumX, sumY int
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(rd, &x, &y)
		sumX += x
		sumY += y
	}

	ans := "Draw"
	if sumX > sumY {
		ans = "Takahashi"
	} else if sumX < sumY {
		ans = "Aoki"
	}

	fmt.Println(ans)
}
