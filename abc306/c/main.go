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

	var a int
	cnt := make([]int, n)
	ans := make([]string, 0, n)
	for i := 0; i < n*3; i++ {
		fmt.Fscan(rd, &a)
		cnt[a-1]++
		if cnt[a-1] == 2 {
			ans = append(ans, strconv.Itoa(a))
		}
	}

	fmt.Fprintln(wr, strings.Join(ans, " "))
}
