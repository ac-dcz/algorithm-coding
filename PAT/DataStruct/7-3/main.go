package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	Val         string
	data        []*Node
	left, right int
}

func (node *Node) Left() *Node {
	if node.left != -1 {
		return node.data[node.left]
	}
	return nil
}

func (node *Node) Right() *Node {
	if node.right != -1 {
		return node.data[node.right]
	}
	return nil
}

func buildTree() (root *Node) {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	nodes, flags := make([]*Node, n), make([]int, n)
	for i := 0; i < n; i++ {
		var v, lf, rf string
		fmt.Fscan(r, &v, &lf, &rf)
		nodes[i] = &Node{Val: v, data: nodes}
		if lf != "-" {
			num, _ := strconv.Atoi(lf)
			flags[num]++
			nodes[i].left = num
		} else {
			nodes[i].left = -1
		}
		if rf != "-" {
			num, _ := strconv.Atoi(rf)
			flags[num]++
			nodes[i].right = num
		} else {
			nodes[i].right = -1
		}
	}
	for i := range flags {
		if flags[i] == 0 {
			root = nodes[i]
			break
		}
	}
	return
}

func Solution(r1, r2 *Node) bool {
	if r1 == nil && r2 == nil {
		return true
	} else if r1 == nil {
		return false
	} else if r2 == nil {
		return false
	} else if r1.Val != r2.Val {
		return false
	} else if r1.left != -1 && r2.left != -1 {
		if r1.Left().Val == r2.Left().Val {
			return Solution(r1.Left(), r2.Left()) && Solution(r1.Right(), r2.Right())
		} else {
			return Solution(r1.Left(), r2.Right()) && Solution(r1.Right(), r2.Left())
		}
	} else if r1.left == -1 && r2.left == -1 {
		return Solution(r1.Right(), r2.Right())
	} else {
		return Solution(r1.Left(), r2.Right()) && Solution(r1.Right(), r2.Left())
	}
}

func main() {
	r1 := buildTree()
	r2 := buildTree()
	ret := Solution(r1, r2)
	if ret {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
