package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n, &m)
	v, w := make([]int, 0), make([]int, 0)
	var a, b, c int
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a, &b, &c)
		for c > 0 {
			v = append(v, a)
			w = append(w, b)
			c--
		}
	}
	dp := make([]int, m+1)
	for i := 0; i < len(v); i++ {
		for j := m; j >= v[i]; j-- {
			if dp[j] < dp[j-v[i]]+w[i] {
				dp[j] = dp[j-v[i]] + w[i]
			}
		}
	}
	fmt.Println(dp[m])
}
