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

	var n, t int
	fmt.Fscan(rd, &n, &t)

	c := make([]int, n)
	r := make([]int, n)

	var et bool
	for i := range c {
		fmt.Fscan(rd, &c[i])
		if c[i] == t {
			et = true
		}
	}

	for i := range r {
		fmt.Fscan(rd, &r[i])
	}

	if !et {
		t = c[0]
	}

	maxR, maxI := -1, -1
	for i := range c {
		if c[i] == t && r[i] > maxR {
			maxR = r[i]
			maxI = i
		}
	}

	fmt.Fprintln(wr, maxI+1)
}
