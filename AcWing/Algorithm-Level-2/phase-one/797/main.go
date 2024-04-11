package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, q int
	fmt.Fscan(r, &n, &q)
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &arr[i])
	}
	for i := n; i > 0; i-- {
		arr[i] -= arr[i-1]
	}
	var x, y, val int
	for q > 0 {
		fmt.Fscan(r, &x, &y, &val)
		arr[x] += val
		if y < n {
			arr[y+1] -= val
		}
		q--
	}
	for i := 1; i <= n; i++ {
		arr[i] += arr[i-1]
		fmt.Printf("%d ", arr[i])
	}
}
