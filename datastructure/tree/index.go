package tree

import "fmt"

func Test() {
	root := CreateTree()
	fmt.Print("前序遍历:")
	PreOrderTraverse(root)
	fmt.Print("\n中序遍历:")
	midOrderTraverse(root)

	fmt.Print("\n二叉查找树插入:")
	te := NewBinarySearchTree(nil)
	te.Insert("30")
	te.Insert(25)
	te.Insert(36)
	te.Insert(20)
	te.Insert(28)
	te.Insert(32)
	te.Insert(40)

	midOrderTraverse(te.root)
	te.Find(28)
	fmt.Printf("\n二叉树查找%v", te.Find(28))
}
