package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1e5 + 1

var parent [N]int

func findRoot(r int) int {
	if parent[r] > 0 {
		parent[r] = findRoot(parent[r])
		return parent[r]
	}
	return r
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m, a, b int
	fmt.Fscan(r, &n, &m)
	for i := 1; i <= n; i++ {
		parent[i] = -1
	}
	var cmd string
	for m > 0 {
		fmt.Fscan(r, &cmd)
		if cmd == "C" {
			fmt.Fscan(r, &a, &b)
			ra, rb := findRoot(a), findRoot(b)
			if ra != rb {
				parent[ra] += parent[rb]
				parent[rb] = ra
			}
		} else if cmd == "Q1" {
			fmt.Fscan(r, &a, &b)
			ra, rb := findRoot(a), findRoot(b)
			if ra == rb {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		} else {
			fmt.Fscan(r, &a)
			ra := findRoot(a)
			fmt.Println(-parent[ra])
		}
		m--
	}
}
