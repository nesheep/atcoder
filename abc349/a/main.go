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

	var sum int
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(rd, &a)
		sum += a
	}

	fmt.Println(-sum)
}
