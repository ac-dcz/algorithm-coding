package main

import (
	"bufio"
	"fmt"
	"os"
)

var pos = [4]int{-3, 3, -1, 1}
var target = 123456789

func main() {
	var (
		val  int
		vt   = make(map[int]int)
		q    [][2]int
		temp = make([]int, 9)
	)
	r := bufio.NewReader(os.Stdin)
	s, _ := r.ReadString('\n')
	for i := range s {
		if ch := s[i]; ch != ' ' && ch != '\r' && ch != '\n' {
			if ch == 'x' {
				val = val*10 + 9
			} else {
				val = val*10 + int(ch-'0')
			}
		}
	}
	vt[val] = 1
	q = append(q, [2]int{val, 0})
	for len(q) > 0 {
		state, d, ind := q[0][0], q[0][1], -1
		if state == target {
			fmt.Println(d)
			break
		}
		for i := 0; i < 9; i++ {
			t := state % 10
			temp[8-i] = t
			if t == 9 {
				ind = 8 - i
			}
			state /= 10
		}
		for i := 0; i < 4; i++ {
			p := ind + pos[i]
			if p < 0 || p > 8 || (i == 2 && ind%3 == 0) || (i == 3 && (ind+1)%3 == 0) {
				continue
			}
			temp[ind], temp[p] = temp[p], temp[ind]
			s := Tonum(temp)
			if _, ok := vt[s]; !ok {
				vt[s] = 1
				q = append(q, [2]int{s, d + 1})
			}
			temp[ind], temp[p] = temp[p], temp[ind]
		}
		q = q[1:]
	}
	if _, ok := vt[target]; !ok {
		fmt.Println(-1)
	}
}

func Tonum(arr []int) int {
	val := 0
	for _, v := range arr {
		val = val*10 + v
	}
	return val
}
