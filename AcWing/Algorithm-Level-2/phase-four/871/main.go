package main

import "fmt"

const mod int = 1e9 + 7

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

	var sum int = 1
	for k, v := range cnts {
		temp, pre := 1, 1
		for i := 0; i < v; i++ {
			temp += pre * k
			pre *= k
			temp, pre = temp%mod, pre%mod
		}
		sum *= temp
		sum %= mod
	}
	fmt.Println(sum)
}
