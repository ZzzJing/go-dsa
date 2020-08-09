package tree

import "fmt"

type Node struct {
	data        interface{}
	left, right *Node
}

func NewNode(v interface{}) *Node {
	return &Node{data: v}
}

func (n *Node) String() string {
	return fmt.Sprintf("v:%+v, left:%+v, right:%+v", n.data, n.left, n.right)
}
