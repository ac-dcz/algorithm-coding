package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n, &m)
	heap := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &heap[i])
	}
	make_heap(heap)
	for m > 0 {
		fmt.Printf("%d ", delete_heap(heap))
		m--
	}
}

func make_heap(heap []int) {
	n := len(heap) - 1
	for i := (n - 1) / 2; i >= 0; i-- {
		update_heap(heap, i)
	}
}

func update_heap(heap []int, parent int) {
	n, val := len(heap), heap[parent]
	for child := 0; 2*parent+1 < n; parent = child {
		child = 2*parent + 1
		if child+1 < n && heap[child+1] < heap[child] {
			child++
		}
		if heap[child] < val {
			heap[parent] = heap[child]
		} else {
			break
		}
	}
	heap[parent] = val
}

func delete_heap(heap []int) int {
	n, val := len(heap), heap[0]
	heap[0] = heap[n-1]
	heap = heap[:n-1]
	update_heap(heap, 0)
	return val
}
