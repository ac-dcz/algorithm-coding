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
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i+1])
	}
	for i := 1; i <= n; i++ {
		arr[i] += arr[i-1]
	}
	var x, y int
	for q > 0 {
		fmt.Fscan(r, &x, &y)
		fmt.Println(arr[y] - arr[x-1])
		q--
	}
}
