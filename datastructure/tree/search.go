package tree

import "errors"

type BinarySearchTree struct {
	root *Node
}

func NewBinarySearchTree(node *Node) *BinarySearchTree {
	return &BinarySearchTree{
		root: node,
	}
}

func (tree *BinarySearchTree) Insert(data interface{}) error {
	var value int
	var ok bool
	if value, ok = data.(int); !ok {
		return errors.New("只支持整形类型")
	}
	if node := tree.root; node == nil {
		tree.root = NewNode(value)
		return nil
	} else {
		//否则找到要插入的位置插入新节点
		for node != nil {
			if value < node.Data.(int) {
				if node.Left == nil {
					node.Left = NewNode(value)
					break
				}
				node = node.Left
			} else if value > node.Data.(int) {
				if node.Right == nil {
					node.Right = NewNode(value)
					break
				}
				node = node.Right
			} else {
				return errors.New("对应数据已经存在")
			}
		}
		return nil
	}
}

//查找
func (tree *BinarySearchTree) Find(data int) *Node {
	if node := tree.root; node == nil {
		return nil
	} else {
		for node != nil {
			if data > node.Data.(int) {
				node = node.Right
			} else if data < node.Data.(int) {
				node = node.Left
			} else {
				return node
			}
		}
		return nil
	}
}
