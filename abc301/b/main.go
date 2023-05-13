package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(r, &n)

	as := make([]int, n)
	blen := n
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &as[i])
		if i > 0 {
			blen += abs(as[i]-as[i-1]) - 1
		}
	}

	bs := make([]string, 0, blen)
	bs = append(bs, strconv.Itoa(as[0]))
	var a, b int
	for i := 1; i < n; i++ {
		a = as[i-1]
		b = as[i]
		if b > a {
			for j := a + 1; j <= b; j++ {
				bs = append(bs, strconv.Itoa(j))
			}
		} else {
			for j := a - 1; j >= b; j-- {
				bs = append(bs, strconv.Itoa(j))
			}
		}
	}

	fmt.Println(strings.Join(bs, " "))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
