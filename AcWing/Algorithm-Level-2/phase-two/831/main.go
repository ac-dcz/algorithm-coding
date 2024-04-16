package main

import (
	"bufio"
	"fmt"
	"os"
)

func MatchNext(str string) []int {
	n := len(str)
	next := make([]int, n)
	next[0] = -1
	// 最大前缀匹配
	for i := 1; i < n; i++ {
		j := next[i-1]
		for j >= 0 && str[j+1] != str[i] {
			j = next[j]
		}
		if str[j+1] == str[i] {
			next[i] = j + 1
		} else {
			next[i] = j
		}
	}
	return next
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	var p, s string
	fmt.Fscan(r, &n, &p, &m, &s)
	next := MatchNext(p)
	i, j := 0, 0
	for i < n && j < m {
		for i < n && j < m {
			if p[i] == s[j] {
				i++
				j++
			} else if i > 0 {
				i = next[i-1] + 1
			} else {
				j++
			}
		}
		if i == n {
			fmt.Printf("%d ", j-n)
			i = next[i-1] + 1
		}
	}
}
