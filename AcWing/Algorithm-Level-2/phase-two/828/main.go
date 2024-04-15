package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	data []int
}

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack) Pop() {
	s.data = s.data[:len(s.data)-1]
}

func (s *Stack) Query() int {
	return s.data[len(s.data)-1]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var (
		cmd    string
		val, n int
	)
	stack := &Stack{data: make([]int, 0)}
	fmt.Fscan(r, &n)
	for n > 0 {
		fmt.Fscan(r, &cmd)
		switch cmd {
		case "push":
			fmt.Fscan(r, &val)
			stack.Push(val)
		case "pop":
			stack.Pop()
		case "query":
			fmt.Println(stack.Query())
		case "empty":
			if stack.Empty() {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
		n--
	}
}
