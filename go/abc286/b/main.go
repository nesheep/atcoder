package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscan(rd, &n, &s)

	ans := strings.Replace(s, "na", "nya", -1)
	fmt.Println(ans)
}
