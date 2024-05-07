package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF int = 0x7f7f7f7f

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m, t int
	fmt.Fscan(r, &n, &m, &t)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			arr[i][j] = INF
		}
		arr[i][i] = 0
	}
	var a, b, c int
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a, &b, &c)
		if arr[a-1][b-1] > c {
			arr[a-1][b-1] = c
		}
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if arr[i][k]+arr[k][j] < arr[i][j] {
					arr[i][j] = arr[i][k] + arr[k][j]
				}
			}
		}
	}

	for t > 0 {
		fmt.Fscan(r, &a, &b)
		if arr[a-1][b-1] >= INF/2 {
			fmt.Println("impossible")
		} else {
			fmt.Println(arr[a-1][b-1])
		}
		t--
	}

}
