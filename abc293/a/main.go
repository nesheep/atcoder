package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var s []byte
	fmt.Fscan(r, &s)

	for i := 0; i < len(s)/2; i++ {
		swap(&s[2*i], &s[2*i+1])
	}

	fmt.Printf("%s\n", s)
}

func swap(a, b *byte) {
	tmp := *a
	*a = *b
	*b = tmp
}
