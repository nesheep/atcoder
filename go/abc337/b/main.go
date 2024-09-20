package main

import (
	"fmt"
	"regexp"
)

func main() {
	var s []byte
	fmt.Scan(&s)

	re := regexp.MustCompile(`^A*B*C*$`)
	if re.Match(s) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
