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

	var n, p, q, r, s int
	fmt.Fscan(rd, &n, &p, &q, &r, &s)

	as := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &as[i])
	}

	bs := make([]string, 0, n)
	for _, v := range as[1:p] {
		bs = append(bs, strconv.Itoa(v))
	}
	for _, v := range as[r : s+1] {
		bs = append(bs, strconv.Itoa(v))
	}
	for _, v := range as[q+1 : r] {
		bs = append(bs, strconv.Itoa(v))
	}
	for _, v := range as[p : q+1] {
		bs = append(bs, strconv.Itoa(v))
	}
	for _, v := range as[s+1:] {
		bs = append(bs, strconv.Itoa(v))
	}

	fmt.Println(strings.Join(bs, " "))
}
