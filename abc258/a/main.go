package main

import (
	"fmt"
	"time"
)

func main() {
	var k int
	fmt.Scan(&k)
	t := time.Date(2000, 1, 1, 21, 0, 0, 0, time.UTC)
	t = t.Add(time.Duration(k) * time.Minute)
	fmt.Println(t.Format("15:04"))
}
