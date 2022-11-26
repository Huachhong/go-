package tree

/*
树遍历
*/

//前序遍历

func PreOrderTraverse(treeNode *Node) {
	//节点为空则退出当前递归
	if treeNode == nil {
		return
	}
	show(treeNode.getData())
	//前序遍历
	PreOrderTraverse(treeNode.Left)
	PreOrderTraverse(treeNode.Right)
}

func midOrderTraverse(treeNode *Node) {
	if treeNode == nil {
		return
	}
	//中序遍历
	midOrderTraverse(treeNode.Left)
	show(treeNode.getData())
	midOrderTraverse(treeNode.Right)
}

func postOrderTraverse(treeNode *Node) {
	if treeNode == nil {
		return
	}
	//后序遍历
	postOrderTraverse(treeNode.Left)
	postOrderTraverse(treeNode.Right)
	show(treeNode.getData())
}
