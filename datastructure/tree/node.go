package tree

import "fmt"

type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func (node *Node) getData() string {
	return fmt.Sprintf("%v", node.Data)
}

func show(data interface{}) {
	fmt.Printf("%v ", data)
}

func CreateTree() *Node {
	node1 := NewNode("A")
	node2 := NewNode("B")
	node3 := NewNode("C")
	node1.Left = node2
	node1.Right = node3

	node21 := NewNode("D")
	node22 := NewNode("E")
	node2.Left = node21
	node2.Right = node22

	node31 := NewNode("F")
	node32 := NewNode("G")
	node3.Left = node31
	node3.Right = node32
	return node1
}
