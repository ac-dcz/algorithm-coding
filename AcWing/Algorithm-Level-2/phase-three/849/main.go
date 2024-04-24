package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const (
	N   int = 1.5e5
	M   int = 1.5e5
	INF int = 0x7f7f7f7f
)

var (
	h       = make([]int, N)
	w       = make([]int, M)
	ne      = make([]int, M)
	nv      = make([]int, M)
	ind int = 0
)

func add(a, b, c int) {
	ne[ind], w[ind], nv[ind] = h[a], c, b
	h[a] = ind
	ind++
}

type IntHeap struct {
	data [][2]int
}

func (h *IntHeap) Len() int {
	return len(h.data)
}

func (h *IntHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *IntHeap) Less(i, j int) bool {
	return h.data[i][1] < h.data[j][1]
}

func (h *IntHeap) Push(val interface{}) {
	h.data = append(h.data, val.([2]int))
}

func (h *IntHeap) Pop() interface{} {
	n := len(h.data)
	x := h.data[n-1]
	h.data = h.data[:n-1]
	return x
}

func dijkstra(n int) int {
	iheap, dist, vt := &IntHeap{make([][2]int, n)}, make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}
	dist[0] = 0
	heap.Push(iheap, [2]int{0, 0})
	for iheap.Len() > 0 {
		temp := heap.Pop(iheap).([2]int)
		t, d := temp[0], temp[1]
		if vt[t] == 1 {
			continue
		}
		vt[t] = 1
		for j := h[t]; j != -1; j = ne[j] {
			v := nv[j]
			if dist[v] > d+w[j] {
				dist[v] = d + w[j]
				heap.Push(iheap, [2]int{v, dist[v]})
			}
		}
	}
	if dist[n-1] == INF {
		return -1
	}
	return dist[n-1]
}

func main() {
	for i := 0; i < N; i++ {
		h[i] = -1
	}
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n, &m)
	var a, b, c int
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a, &b, &c)
		add(a-1, b-1, c)
	}
	fmt.Println(dijkstra(n))
}
