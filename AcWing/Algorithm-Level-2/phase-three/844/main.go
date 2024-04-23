package main

import (
	"bufio"
	"fmt"
	"os"
)

var pos = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
var n, m int

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)
	arr, vt := make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i], vt[i] = make([]int, m), make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(r, &arr[i][j])
		}
	}
	q := [][3]int{{0, 0, 0}}
	vt[0][0] = 1
	for len(q) > 0 {
		x, y, d := q[0][0], q[0][1], q[0][2]
		if x == n-1 && y == m-1 {
			fmt.Println(d)
			break
		}
		for i := 0; i < 4; i++ {
			px, py := x+pos[i][0], y+pos[i][1]
			if px < 0 || px >= n || py < 0 || py >= m || arr[px][py] == 1 || vt[px][py] == 1 {
				continue
			}
			vt[px][py] = 1
			q = append(q, [3]int{px, py, d + 1})
		}
		q = q[1:]
	}
}
