package main

import "fmt"

type TrNode struct {
	Val   int
	Left  *TrNode
	Right *TrNode
}

func (t *TrNode) InsertNode(node *TrNode) *TrNode {
	if t == nil {
		return node
	}
	if t.Val < node.Val {
		t.Right = t.Right.InsertNode(node)
	} else if t.Val > node.Val {
		t.Left = t.Left.InsertNode(node)
	}
	return t
}

func (t *TrNode) Equal(n *TrNode) bool {
	if t == nil && n == nil {
		return true
	} else if t == nil || n == nil {
		return false
	} else if t.Val != n.Val {
		return false
	} else {
		return t.Left.Equal(n.Left) && t.Right.Equal(n.Right)
	}
}

type Tree struct {
	root *TrNode
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) InsertNode(node *TrNode) {
	t.root = t.root.InsertNode(node)
}

func (t *Tree) Equal(o *Tree) bool {
	return t.root.Equal(o.root)
}

func main() {
	var n, v, w int
	for {
		if fmt.Scan(&n); n == 0 {
			break
		}
		fmt.Scan(&v)
		root := NewTree()
		for i := 0; i < n; i++ {
			fmt.Scan(&w)
			root.InsertNode(&TrNode{Val: w})
		}
		for v > 0 {
			t := NewTree()
			for i := 0; i < n; i++ {
				fmt.Scan(&w)
				t.InsertNode(&TrNode{Val: w})
			}
			if root.Equal(t) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
			v--
		}
	}
}
