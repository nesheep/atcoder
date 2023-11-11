package main

import "fmt"

func main() {
	var s []byte
	fmt.Scan(&s)

	t := make([]bool, len(s))
	i := 0
	for i < len(s)-2 {
		if t[i] || s[i] != 'A' {
			i++
			continue
		}
		j := i + 1
		for j < len(s)-1 {
			if t[j] {
				j++
				continue
			}
			if s[j] != 'B' {
				i++
				break
			}
			k := j + 1
			for k < len(s) {
				if t[k] {
					k++
					continue
				}
				if s[k] != 'C' {
					i++
					break
				}
				t[i] = true
				t[j] = true
				t[k] = true
				var cnt int
				for l := i - 1; l >= 0; l-- {
					if t[l] {
						continue
					}
					cnt++
					if cnt == 2 {
						i = l
						break
					}
				}
				break
			}
			break
		}
	}

	ans := make([]byte, 0, len(s))
	for i := range s {
		if t[i] {
			continue
		}
		ans = append(ans, s[i])
	}

	fmt.Printf("%s\n", ans)
}
