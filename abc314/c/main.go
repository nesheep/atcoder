package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, m int
	var s []byte
	fmt.Fscan(rd, &n, &m, &s)

	sColor := make([][2]int, 0, n)
	colorPos := make(map[int][]byte, m)
	for i := 0; i < n; i++ {
		var c int
		fmt.Fscan(rd, &c)
		sColor = append(sColor, [2]int{c, len(colorPos[c])})
		colorPos[c] = append(colorPos[c], s[i])
	}

	t := make([]byte, 0, n)
	for _, sc := range sColor {
		c := sc[0]
		i := sc[1] - 1
		if i < 0 {
			i = len(colorPos[c]) - 1
		}
		t = append(t, colorPos[c][i])
	}

	fmt.Fprintf(wr, "%s\n", t)
}
