package main

import (
	"bufio"
	"fmt"
	"os"
)

var bucket [51]int

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, v int
	fmt.Fscan(r, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &v)
		bucket[v]++
	}
	for i, val := range bucket {
		if val > 0 {
			fmt.Printf("%d:%d\n", i, val)
		}
	}
}
