package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, v int
	fmt.Fscan(r, &n)
	for n > 0 {
		fmt.Fscan(r, &v)
		p1, p2 := make([]int, 0), make([]int, 0)
		for i := 1; i*i <= v; i++ {
			if v%i == 0 {
				p1 = append(p1, i)
				if i*i != v {
					p2 = append(p2, v/i)
				}
			}
		}
		for _, v := range p1 {
			fmt.Printf("%d ", v)
		}
		for i := len(p2) - 1; i >= 0; i-- {
			fmt.Printf("%d ", p2[i])
		}
		fmt.Println()
		n--
	}
}
