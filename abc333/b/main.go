package main

import "fmt"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	var s, t []byte
	fmt.Scan(&s, &t)

	ds := abs(int(s[0]) - int(s[1]))
	dt := abs(int(t[0]) - int(t[1]))
	if ds == dt || ds == 5-dt {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
