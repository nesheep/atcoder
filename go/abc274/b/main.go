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

	var h, w int
	fmt.Fscan(rd, &h, &w)

	c := make([][]byte, h)
	for i := range c {
		fmt.Fscan(rd, &c[i])
	}

	cnt := make([]int, w)
	for i := range c {
		for j := range c[i] {
			if c[i][j] == '#' {
				cnt[j]++
			}
		}
	}

	ans := make([]string, 0, w)
	for _, v := range cnt {
		ans = append(ans, strconv.Itoa(v))
	}

	fmt.Fprintln(wr, strings.Join(ans, " "))
}
