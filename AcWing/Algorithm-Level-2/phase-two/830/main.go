package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, val int
	fmt.Fscan(r, &n)
	var stack []int
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &val)
		for len(stack) > 0 && stack[len(stack)-1] >= val {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			fmt.Printf("%d ", -1)
		} else {
			fmt.Printf("%d ", stack[len(stack)-1])
		}
		stack = append(stack, val)
	}

}
