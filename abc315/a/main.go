package main

import "fmt"

func main() {
	var s []byte
	fmt.Scan(&s)

	t := make([]byte, 0, len(s))
	for _, c := range s {
		if c == 'a' || c == 'i' || c == 'u' || c == 'e' || c == 'o' {
			continue
		}
		t = append(t, c)
	}

	fmt.Println(string(t))
}
