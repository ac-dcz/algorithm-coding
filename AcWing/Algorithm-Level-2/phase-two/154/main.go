package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, k, val int
	fmt.Fscan(r, &n, &k)
	var minArr, maxArr, qmin, qmax []int
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &val)
		if i == k {
			qmin = qmin[1:]
			qmax = qmax[1:]
		}
		for len(qmin) > 0 && qmin[len(qmin)-1] > val {
			qmin = qmin[:len(qmin)-1]
		}
		for len(qmax) > 0 && qmax[len(qmax)-1] < val {
			qmax = qmax[:len(qmax)-1]
		}
		qmin = append(qmin, val)
		qmax = append(qmax, val)
		if i+1 >= k {
			minArr = append(minArr, qmin[0])
			maxArr = append(maxArr, qmax[0])
		}
	}
	for _, val := range minArr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
	for _, val := range maxArr {
		fmt.Printf("%d ", val)
	}
}
