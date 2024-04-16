package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1e5 + 1

var parent [N]int

func findRoot(r int) int {
	if parent[r] != 0 {
		parent[r] = findRoot(parent[r])
		return parent[r]
	}
	return r
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, a, b int
	var cmd string
	fmt.Fscan(r, &n, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &cmd, &a, &b)
		if cmd == "M" {
			ra, rb := findRoot(a), findRoot(b)
			if ra != rb {
				parent[ra] = rb
			}
		} else {
			ra, rb := findRoot(a), findRoot(b)
			if ra == rb {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}
