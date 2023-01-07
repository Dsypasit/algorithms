package mergetwolist

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func newNode(v int, node *Node) *Node {
	return &Node{
		Val:  v,
		Next: node,
	}
}

func addAll(current *Node, nodes ...*Node) *Node {
	if len(nodes) == 0 {
		return current
	}
	current.Next = nodes[0]
	addAll(nodes[0], nodes[1:]...)
	return current
}

func (n *Node) print() {
	fmt.Println(n.Val)
	if n.Next == nil {
		return
	}
	n.Next.print()
}

func (n *Node) add(node *Node) {
	n.Next = node
}

func equal(l1, l2 *Node) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	if l1 == nil && l2 == nil {
		return false
	}
	if l1.Val != l2.Val {
		return false
	}
	return equal(l1.Next, l2.Next)
}
