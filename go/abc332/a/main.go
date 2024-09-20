package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, s, k int
	fmt.Fscan(rd, &n, &s, &k)

	var sum int

	for i := 0; i < n; i++ {
		var p, q int
		fmt.Fscan(rd, &p, &q)
		sum += p * q
	}

	if sum < s {
		sum += k
	}

	fmt.Println(sum)
}
