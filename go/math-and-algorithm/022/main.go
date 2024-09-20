package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	cnt := make([]int, 100000)
	for _, v := range a {
		cnt[v]++
	}

	var ans int
	for i := 1; i < 50000; i++ {
		ans += cnt[i] * cnt[100000-i]
	}

	ans += (cnt[50000] * (cnt[50000] - 1)) / 2

	fmt.Fprintln(w, ans)
}
