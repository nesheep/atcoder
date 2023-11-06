package main

import "fmt"

func main() {
	var s []byte
	fmt.Scan(&s)

	for i := 0; i < 6; i += len(s) {
		fmt.Printf("%s", s)
	}
	fmt.Print("\n")
}
