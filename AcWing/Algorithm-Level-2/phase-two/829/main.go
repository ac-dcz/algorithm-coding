package main

import (
	"bufio"
	"fmt"
	"os"
)

type Queue struct {
	data []int
}

func (q *Queue) Push(val int) {
	q.data = append(q.data, val)
}

func (q *Queue) Pop() {
	q.data = q.data[1:]
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

func (q *Queue) Query() int {
	return q.data[0]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var (
		n, val int
		cmd    string
	)
	q := &Queue{data: make([]int, 0)}
	fmt.Fscan(r, &n)
	for n > 0 {
		fmt.Fscan(r, &cmd)
		switch cmd {
		case "push":
			fmt.Fscan(r, &val)
			q.Push(val)
		case "pop":
			q.Pop()
		case "query":
			fmt.Println(q.Query())
		case "empty":
			if q.Empty() {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
		n--
	}
}
