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
	fmt.Fscan(rd, &n, &m)

	c := make([]string, n)
	for i := range c {
		fmt.Fscan(rd, &c[i])
	}

	d := make([]string, m)
	for i := range d {
		fmt.Fscan(rd, &d[i])
	}

	var p0 int
	fmt.Fscan(rd, &p0)

	p := make([]int, m)
	for i := range p {
		fmt.Fscan(rd, &p[i])
	}

	dpm := make(map[string]int, m)
	for i := range d {
		dpm[d[i]] = p[i]
	}

	var ans int
	for i := range c {
		if v, ok := dpm[c[i]]; ok {
			ans += v
		} else {
			ans += p0
		}
	}

	fmt.Fprintln(wr, ans)
}
