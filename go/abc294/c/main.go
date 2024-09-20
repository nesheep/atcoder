package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, m int
	fmt.Fscan(rd, &n, &m)

	a := make([]int, n)
	b := make([]int, m)
	c := make([]int, 0, n+m)
	for i := range a {
		fmt.Fscan(rd, &a[i])
		c = append(c, a[i])
	}
	for i := range b {
		fmt.Fscan(rd, &b[i])
		c = append(c, b[i])
	}

	sort.Ints(c)

	cm := make(map[int]int, len(c))
	for i, v := range c {
		cm[v] = i
	}

	ansA := make([]string, n)
	ansB := make([]string, m)
	for i := range a {
		ansA[i] = strconv.Itoa(cm[a[i]] + 1)
	}
	for i := range b {
		ansB[i] = strconv.Itoa(cm[b[i]] + 1)
	}

	fmt.Fprintln(wr, strings.Join(ansA, " "))
	fmt.Fprintln(wr, strings.Join(ansB, " "))
}
