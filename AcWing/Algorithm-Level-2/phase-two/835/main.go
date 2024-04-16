package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1e5 + 1

var (
	tire [N][26]int
	cnts [N]int
	ind  int = 0
)

func Insert(str string) {
	p := 0
	for i := range str {
		j := int(str[i] - 'a')
		if tire[p][j] == 0 {
			tire[p][j] = ind + 1
			ind++
		}
		p = tire[p][j]
	}
	cnts[p]++
}

func Query(str string) int {
	p := 0
	for i := range str {
		j := int(str[i] - 'a')
		if tire[p][j] == 0 {
			return 0
		}
		p = tire[p][j]
	}
	return cnts[p]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	var cmd, val string
	fmt.Fscan(r, &n)
	for n > 0 {
		fmt.Fscan(r, &cmd, &val)
		if cmd == "I" {
			Insert(val)
		} else {
			fmt.Println(Query(val))
		}
		n--
	}
}
