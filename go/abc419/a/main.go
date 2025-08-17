package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(engToAC(s))
}

func engToAC(s string) string {
	switch s {
	case "red":
		return "SSS"
	case "blue":
		return "FFF"
	case "green":
		return "MMM"
	default:
		return "Unknown"
	}
}
