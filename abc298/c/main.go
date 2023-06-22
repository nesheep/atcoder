package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func answer(s []int) {
	sort.Ints(s)
	t := make([]string, 0, len(s))
	for _, v := range s {
		t = append(t, strconv.Itoa(v))
	}
	fmt.Println(strings.Join(t, " "))
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, q int
	fmt.Fscan(rd, &n, &q)

	a := make(map[int][]int)
	b := make(map[int]map[int]bool)

	var qq, qi, qj int
	for i := 0; i < q; i++ {
		fmt.Fscan(rd, &qq, &qi)
		switch qq {
		case 1:
			fmt.Fscan(rd, &qj)
			a[qj] = append(a[qj], qi)
			if b[qi] == nil {
				b[qi] = make(map[int]bool)
			}
			b[qi][qj] = true
		case 2:
			answer(a[qi])
		case 3:
			s := make([]int, 0, len(b[qi]))
			for k := range b[qi] {
				s = append(s, k)
			}
			answer(s)
		}
	}
}
