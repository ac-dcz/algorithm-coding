package main

import (
	"fmt"
)

func main() {
	cnts := make(map[int]int)
	var n, v int
	fmt.Scan(&n)
	for n > 0 {
		fmt.Scan(&v)
		for i := 2; i*i <= v; i++ {
			ret := 0
			for v%i == 0 {
				v /= i
				ret++
			}
			if ret > 0 {
				cnts[i] += ret
			}
		}
		if v != 1 {
			cnts[v]++
		}
		n--
	}
	var (
		sum int = 1
		mod int = 1e9 + 7
	)
	for _, v := range cnts {
		sum *= v + 1
		sum %= mod
	}
	fmt.Println(sum)
}
