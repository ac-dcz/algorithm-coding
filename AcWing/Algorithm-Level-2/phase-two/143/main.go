package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 31e5 + 1

var (
	tire [N][2]int
	ind  int = 0
)

func insert(num int) int {
	p, t, ret := 0, 0, 0
	for i := 30; i >= 0; i-- {
		b := (num >> i) & 1
		if tire[p][b] == 0 {
			tire[p][b] = ind + 1
			ind++
		}
		if tire[t][b^1] != 0 {
			t = tire[t][b^1]
			ret = ret*2 + (b ^ 1)
		} else {
			t = tire[t][b]
			ret = ret*2 + b
		}
		p = tire[p][b]
	}
	return ret
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, val, ant int
	fmt.Fscan(r, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &val)
		ret := insert(val)
		if ret^val > ant {
			ant = ret ^ val
		}
	}
	fmt.Println(ant)
}
