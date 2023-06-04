package main

import (
	"bufio"
	"fmt"
	"os"
)

type person struct {
	name string
	age  int
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	ps := make([]person, n)
	for i := range ps {
		fmt.Fscan(rd, &ps[i].name, &ps[i].age)
	}

	min := ps[0].age
	var pmin int
	for i := 1; i < n; i++ {
		if ps[i].age < min {
			min = ps[i].age
			pmin = i
		}
	}

	for _, p := range ps[pmin:] {
		fmt.Fprintln(wr, p.name)
	}
	for _, p := range ps[:pmin] {
		fmt.Fprintln(wr, p.name)
	}
}
