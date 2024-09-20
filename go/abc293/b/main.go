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

	as := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &as[i])
	}

	bs := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		if !bs[i] {
			bs[as[i]] = true
		}
	}

	xs := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		if !bs[i] {
			xs = append(xs, strconv.Itoa(i))
		}
	}

	fmt.Println(len(xs))
	fmt.Println(strings.Join(xs, " "))
}
