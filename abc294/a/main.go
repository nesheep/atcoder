package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(r, &n)

	var a int
	bs := make([]string, 0, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a)
		if a%2 == 0 {
			bs = append(bs, strconv.Itoa(a))
		}
	}

	fmt.Println(strings.Join(bs, " "))
}
