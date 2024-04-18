package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1e5 + 1

var (
	heap  = make([][2]int, N)
	index = make([]int, 0) //第几个插入的，在第几个位置
	ind   = 0
)

func insert_up(i int) {
	val, p := heap[i][0], heap[i][1]
	for ; i > 0 && heap[(i-1)/2][0] > val; i = (i - 1) / 2 {
		heap[i] = heap[(i-1)/2]
		index[heap[i][1]] = i
	}
	heap[i] = [2]int{val, p}
	index[p] = i
}

func insert_down(i int) {
	temp := heap[i]
	for j := 0; 2*i+1 < ind; i = j {
		j = 2*i + 1
		if j < ind && heap[j+1][0] < heap[j][0] {
			j++
		}
		if heap[j][0] < temp[0] {
			heap[i] = heap[j]
			index[heap[i][1]] = i
		} else {
			break
		}
	}
	heap[i] = temp
	index[temp[1]] = i
}

func delete_heap(k int) {
	if k == -1 {
		heap[0] = heap[ind-1]
		index[heap[0][1]] = 0
		ind--
		insert_down(0)
	} else {
		i := index[k]
		heap[i] = heap[ind-1]
		index[heap[i][1]] = i
		ind--
		insert_down(i)
		insert_up(i)
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, k, val int
	var cmd string
	fmt.Fscan(r, &n)
	for n > 0 {
		fmt.Fscan(r, &cmd)
		switch cmd {
		case "I":
			fmt.Fscan(r, &val)
			heap[ind] = [2]int{val, len(index)}
			index = append(index, ind)
			insert_up(ind)
			ind++
		case "PM":
			fmt.Println(heap[0][0])
		case "DM":
			delete_heap(-1)
		case "D":
			fmt.Fscan(r, &k)
			delete_heap(k - 1)
		case "C":
			fmt.Fscan(r, &k, &val)
			i := index[k-1]
			heap[i][0] = val
			insert_down(i)
			insert_up(i)
		}
		n--
	}
}
