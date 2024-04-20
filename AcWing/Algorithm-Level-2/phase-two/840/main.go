package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 4e5 + 3

var arr = make([]*int, N)

func hashKey(val int) int {
	if val < 0 {
		val = -val
	}
	return val % N
}

func findIndex(num int) int {
	key := hashKey(num)
	for arr[key] != nil && (*arr[key]) != num {
		key = (key + 1) % N
	}
	return key
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var (
		cmd  string
		x, n int
	)
	fmt.Fscan(r, &n)
	for n > 0 {
		fmt.Fscan(r, &cmd, &x)
		ind := findIndex(x)
		if cmd == "I" {
			val := x
			arr[ind] = &val
		} else {
			if arr[ind] != nil {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
		n--
	}

}
