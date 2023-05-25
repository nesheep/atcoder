package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	r := bufio.NewReader(os.Stdin)

	var n int
	var s []byte
	fmt.Fscan(r, &n, &s)

	p := point{}
	m := make(map[point]bool, n)
	var ans bool

	for _, v := range s {
		m[p] = true
		switch v {
		case 'R':
			p.x++
		case 'L':
			p.x--
		case 'U':
			p.y++
		case 'D':
			p.y--
		}
		if m[p] {
			ans = true
			break
		}
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
