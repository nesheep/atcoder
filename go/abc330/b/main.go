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

	var n, l, r int
	fmt.Fscan(rd, &n, &l, &r)

	ans := make([]string, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(rd, &a)
		if a < l {
			ans[i] = strconv.Itoa(l)
		} else if a > r {
			ans[i] = strconv.Itoa(r)
		} else {
			ans[i] = strconv.Itoa(a)
		}
	}

	fmt.Println(strings.Join(ans, " "))
}
