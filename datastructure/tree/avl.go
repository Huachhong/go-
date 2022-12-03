package tree

//	平衡二叉树节点类
type AVLNode struct {
	Data        int
	Left, Right *AVLNode
	Height      int
}

//	平衡二叉树结构体
type AVLTree struct {
	root *AVLNode
}

func NewAVLTree(data int) *AVLTree {
	return &AVLTree{
		root: &AVLNode{Data: data, Height: 1},
	}
}

func (tree *AVLTree) Find(data int) *AVLNode {
	if tree.root == nil {
		return nil
	}
	return tree.root.Find(data)
}

//	查找指定节点
func (node *AVLNode) Find(data int) *AVLNode {
	if node.Data == data {
		return node
	} else if data < node.Data {
		//	如果查找的值小于节点值，从节点的左子树开始查找
		if node.Left == nil {
			return nil
		}
		return node.Left.Find(data)
	} else {
		//	如果查找的值大于节点值，从节点的右子树开始查找
		if node.Right == nil {
			return nil
		}
		return node.Right.Find(data)
	}
}
