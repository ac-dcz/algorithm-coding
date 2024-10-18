package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Heap struct {
	data []int
}

func NewHeap() *Heap {
	return &Heap{}
}

func (h *Heap) Insert(val int) {
	h.data = append(h.data, val)
	i := len(h.data) - 1
	for ; i != 0 && h.data[(i-1)/2] > val; i = (i - 1) / 2 {
		h.data[i] = h.data[(i-1)/2]
	}
	h.data[i] = val
}

func (h *Heap) Path(ind int) []string {
	var items []string
	for {
		items = append(items, strconv.Itoa(h.data[ind]))
		if ind == 0 {
			break
		}
		ind = (ind - 1) / 2
	}
	return items
}

func main() {
	var n, m, v int
	fmt.Scan(&n, &m)
	heap := NewHeap()
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		heap.Insert(v)
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&v)
		path := heap.Path(v - 1)
		fmt.Println(strings.Join(path, " "))
	}
}
