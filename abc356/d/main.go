package main

import "fmt"

const d = 998244353

func popcount(n uint64) uint64 {
	var ans uint64
	for n > 0 {
		ans += n & 1
		n >>= 1
	}
	return ans
}

func main() {
	var n, m uint64
	fmt.Scan(&n, &m)

	var ans uint64
	for i := uint64(0); i < n; i++ {
		ans += popcount(i & m)
		ans %= d
	}

	fmt.Println(ans)
}
