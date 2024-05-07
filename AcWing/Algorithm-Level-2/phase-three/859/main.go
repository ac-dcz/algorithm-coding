package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const N int = 1e5

type HeapInt struct {
	data [][3]int
}

func (h *HeapInt) Len() int {
	return len(h.data)
}

func (h *HeapInt) Less(i, j int) bool {
	return h.data[i][2] < h.data[j][2]
}

func (h *HeapInt) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *HeapInt) Push(val interface{}) {
	v := val.([3]int)
	h.data = append(h.data, v)
}

func (h *HeapInt) Pop() interface{} {
	n := len(h.data)
	val := h.data[n-1]
	h.data = h.data[:n-1]
	return val
}

var parent = make([]int, N)

func findroot(r int) int {
	if parent[r] >= 0 {
		parent[r] = findroot(parent[r])
		return parent[r]
	}
	return r
}

func main() {
	for i := range parent {
		parent[i] = -1
	}
	h := &HeapInt{}
	r := bufio.NewReader(os.Stdin)
	var n, m, a, b, c int
	fmt.Fscan(r, &n, &m)
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a, &b, &c)
		heap.Push(h, [3]int{a - 1, b - 1, c})
	}
	ret := 0
	for h.Len() > 0 && n > 0 {
		temp := heap.Pop(h).([3]int)
		a, b, c = temp[0], temp[1], temp[2]
		ra, rb := findroot(a), findroot(b)
		if ra == rb {
			continue
		}
		ret += c
		parent[ra] = rb
		n--
	}
	if n != 1 {
		fmt.Println("impossible")
	} else {
		fmt.Println(ret)
	}
}
