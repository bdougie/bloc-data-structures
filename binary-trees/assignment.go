package main

import (
	"fmt"
	// "math/rand"
)

type Node struct {
	Right, Left, parent *Node
	value               int
}

type BinaryTree struct {
	root *Node
}

func addNode(parent *Node, v int) *Node {
	return &Node{nil, nil, parent, v}
}

func (b *BinaryTree) insert(root *Node, parent *Node, v int) *Node {
	switch {
	case root == nil:
		return addNode(parent, v)
	case v <= root.value:
		root.Left = b.insert(root.Left, root, v)
	case v > root.value:
		root.Right = b.insert(root.Right, root, v)
	}
	return root
}

func (b *BinaryTree) Insert(val int) (n *Node) {
	if b.root == nil {
		n = addNode(nil, val)
		b.root = n
	} else {
		n = b.insert(b.root, nil, val)
	}
	return
}

func (n *Node) Print() string {
	if n == nil {
		return "()"
	}
	s := ""
	if n.Left != nil {
		s += n.Left.Print() + " "
	}
	s += fmt.Sprint(n.value)
	if n.Right != nil {
		s += " " + n.Right.Print()
	}
	return "(" + s + ")"
}

func (b *BinaryTree) Delete(val int) {
	if b.root == nil {
		return
	} else {
		b.delete(b.root, val)
	}
}

func (b *BinaryTree) delete(root *Node, v int) (n *Node) {
	if b.root.value == v && b.root.Left == nil && b.root.Right == nil {
		n = b.root
		b.root = nil
		return
	} else if root.value != v {
		if v <= root.value {
			b.delete(root.Left, v)
		} else if v > root.value {
			b.delete(root.Right, v)
		}
	} else if root.value == v && root.Left == nil && root.Right == nil {
		n = b.root
		b.root = nil
		return
	} else if root.value == v && (root.Left == nil || root.Right == nil) {
		// there's a better, cleaner way to do this in one line
		var n1 *Node
		if root.Left != nil {
			n1 = root.Left
		} else {
			n1 = root.Right
		}
		switch {
		case root.parent.Left == root:
			root.parent.Left = n1
			return root
		case root.parent.Right == root:
			root.parent.Right = n1
			return root
		}
	} else if root.value == v && root.Left != nil && root.Right != nil {
		// Needs to be implemented later
	}
	return nil
}

func main() {}
