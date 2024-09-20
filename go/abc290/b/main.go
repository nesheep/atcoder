package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n, k int
	var s []byte
	fmt.Fscan(r, &n, &k, &s)

	ans := make([]byte, n)
	var tmp int
	for i, v := range s {
		if tmp < k && v == 'o' {
			ans[i] = 'o'
			tmp++
		} else {
			ans[i] = 'x'
		}
	}

	fmt.Printf("%s\n", ans)
}
