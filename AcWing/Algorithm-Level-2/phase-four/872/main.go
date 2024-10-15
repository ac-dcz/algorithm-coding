package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, a, b int
	fmt.Fscan(r, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a, &b)
		fmt.Println(gcd(a, b))
	}
}
