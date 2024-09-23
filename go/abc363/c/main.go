package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	var s []byte
	fmt.Scan(&n, &k, &s)

	var ans int
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	for {
		var f bool
		for i := 0; i <= n-k; i++ {
			sub := s[i : i+k]
			if isPalindrome(sub) {
				f = true
				break
			}
		}

		if !f {
			ans++
		}

		if !nextPermutation(s) {
			break
		}
	}

	fmt.Println(ans)
}

func nextPermutation(s []byte) bool {
	step1 := func(s []byte) int {
		k := -1
		for i := len(s) - 2; i >= 0; i-- {
			if s[i] < s[i+1] {
				k = i
				break
			}
		}
		return k
	}

	step2 := func(s []byte, k int) int {
		l := -1
		for i := len(s) - 1; i > k; i-- {
			if s[k] < s[i] {
				l = i
				break
			}
		}
		return l
	}

	step3 := func(s []byte, k, l int) {
		s[k], s[l] = s[l], s[k]
	}

	step4 := func(s []byte, k int) {
		for i, j := k+1, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}

	k := step1(s)
	if k < 0 {
		return false
	}
	l := step2(s, k)
	step3(s, k, l)
	step4(s, k)

	return true
}

func isPalindrome(s []byte) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
