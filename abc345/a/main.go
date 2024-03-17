package main

import (
	"fmt"
	"regexp"
)

var re = regexp.MustCompile(`^<=+>$`)

func main() {
	var s string
	fmt.Scan(&s)
	if re.MatchString(s) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
