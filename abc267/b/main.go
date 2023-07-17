package main

import "fmt"

func judge(t [10]bool) bool {
	if t[0] {
		return false
	}
	var l [7]bool
	l[0] = t[6]
	l[1] = t[3]
	l[2] = t[1] || t[7]
	l[3] = t[4]
	l[4] = t[2] || t[8]
	l[5] = t[5]
	l[6] = t[9]
	for i := 0; i < 7; i++ {
		if !l[i] {
			continue
		}
		for j := i + 2; j < 7; j++ {
			if !l[j] {
				continue
			}
			f := true
			for k := i + 1; k < j; k++ {
				if l[k] {
					f = false
					break
				}
			}
			if f {
				return true
			}
		}
	}
	return false
}

func main() {
	var s []byte
	fmt.Scan(&s)

	var t [10]bool
	for i := range t {
		t[i] = s[i] == '1'
	}

	ans := "No"
	if judge(t) {
		ans = "Yes"
	}

	fmt.Println(ans)
}
