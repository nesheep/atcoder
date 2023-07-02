package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type persion struct {
	id, a, b int
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	p := make([]persion, n)
	for i := range p {
		p[i].id = i + 1
		fmt.Fscan(rd, &p[i].a, &p[i].b)
	}

	sort.Slice(p, func(i, j int) bool {
		d := p[i].a*(p[j].a+p[j].b) - p[j].a*(p[i].a+p[i].b)
		if d == 0 {
			return p[i].id < p[j].id
		}
		return d > 0
	})

	ans := make([]string, 0, len(p))
	for i := range p {
		ans = append(ans, strconv.Itoa(p[i].id))
	}

	fmt.Fprintln(wr, strings.Join(ans, " "))
}
