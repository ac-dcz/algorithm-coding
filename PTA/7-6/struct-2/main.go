package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dfs(graph [][]int, v int, vt []int) []string {
	vt[v] = 1
	temp := []string{strconv.Itoa(v)}
	for i, w := range graph[v] {
		if w == 1 && vt[i] == 0 {
			temp = append(temp, dfs(graph, i, vt)...)
		}
	}
	return temp
}

func bfs(graph [][]int, v int, vt []int) []string {
	var q []int
	var temp []string
	vt[v] = 1
	q = append(q, v)
	for len(q) > 0 {
		v := q[0]
		temp = append(temp, strconv.Itoa(v))
		for i, w := range graph[v] {
			if w == 1 && vt[i] == 0 {
				vt[i] = 1
				q = append(q, i)
			}
		}
		q = q[1:]
	}
	return temp
}

func main() {
	var graph [][]int
	var n, m, a, b int
	fmt.Scan(&n, &m)
	graph = make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&a, &b)
		graph[a][b], graph[b][a] = 1, 1
	}
	vt := make([]int, n)
	for i := 0; i < n; i++ {
		if vt[i] == 0 {
			data := dfs(graph, i, vt)
			fmt.Printf("{ %s }\n", strings.Join(data, " "))
		}
	}
	vt = make([]int, n)
	for i := 0; i < n; i++ {
		if vt[i] == 0 {
			data := bfs(graph, i, vt)
			fmt.Printf("{ %s }\n", strings.Join(data, " "))
		}
	}

}
