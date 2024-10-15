package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	vt := make([]int, n+1)
	ret := 0
	for i := 2; i <= n; i++ {
		if vt[i] == 0 {
			for k := 2; k*i <= n; k++ {
				vt[k*i] = 1
			}
			ret++
		}
	}
	fmt.Println(ret)
}
