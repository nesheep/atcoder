package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	a := make([]int, n)
	var v int
	for i := 0; i < n*7; i++ {
		fmt.Fscan(rd, &v)
		a[i/7] += v
	}

	ans := make([]string, 0, n)
	for i := range a {
		ans = append(ans, strconv.Itoa(a[i]))
	}

	fmt.Fprintln(wr, strings.Join(ans, " "))
}
