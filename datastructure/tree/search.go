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

func (tree *BinarySearchTree) Delete(data int) error {
	if tree.root == nil {
		return errors.New("二叉查找树为空")
	}
	p := tree.root
	var pp *Node = nil
	//	找到待删除节点
	for p != nil && p.Data.(int) != data {
		pp = p
		if p.Data.(int) < data {
			p = p.Right
		} else {
			p = p.Left
		}
	}
	if p == nil {
		return errors.New("待删除节点不存在")
	}
	if p.Left != nil && p.Right != nil {
		minP := p.Right //	右子树中的最小节点
		minPP := p      //	minP的父节点
		//	查找右子树中的最小节点
		for minP.Left != nil {
			minPP = minP
			minP = p.Left
		}
		p.Data = minP.Data // 将minP 的数据设置到p中
		p = minP           // 下面就变成删除minP
		pp = minPP
	}
	//	应用待删除节点只有左右节点的逻辑
	var child *Node
	if p.Left != nil {
		child = p.Left
	} else if p.Right != nil {
		child = p.Right
	} else {
		child = nil
	}
	if pp == nil {
		//	删除根节点
		tree.root = nil
	} else if pp.Left == p {
		//	仅有左节点
		pp.Left = child
	} else if pp.Right == p {
		//	仅有右节点
		pp.Right = child
	}
	return nil
}
