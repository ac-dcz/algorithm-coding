package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i])
		m := 0
		for arr[i] > 0 {
			arr[i] &= arr[i] - 1
			m++
		}
		fmt.Printf("%d ", m)
	}

}
