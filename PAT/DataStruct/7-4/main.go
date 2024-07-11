package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (root *Node) Insert(node *Node) *Node {
	if root == nil {
		return node
	} else {
		if root.Val < node.Val {
			root.Right = node.Right.Insert(node)
		} else if root.Val >= node.Val {
			root.Left = node.Left.Insert(node)
		}
		return root
	}
}

func (root *Node) Travle(val int, flags []int) bool {
	if root == nil {
		return false
	}
	if root.Val != val {
		if flags[root.Val] == 1 {
			return false
		} else if root.Val < val {
			return root.Right.Travle(val, flags)
		} else {
			return root.Left.Travle(val, flags)
		}
	}
	flags[val] = 1
	return true
}

type Tree struct {
	root *Node
}

func NewTree(arr []int) *Tree {
	tr := &Tree{}
	for _, val := range arr {
		node := &Node{Val: val}
		tr.Insert(node)
	}
	return tr
}

func (tr *Tree) Insert(node *Node) {
	tr.root = tr.root.Insert(node)
}

func (tr *Tree) Check(arr []int) bool {
	flag := make([]int, len(arr)+1)
	for _, val := range arr {
		if !tr.root.Travle(val, flag) {
			return false
		}
	}
	return true
}

func main() {
	var n, m int
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		fmt.Scan(&m)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}
		tr := NewTree(arr)
		for m > 0 {
			for i := 0; i < n; i++ {
				fmt.Scan(&arr[i])
			}
			if tr.Check(arr) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
			m--
		}
	}
}
