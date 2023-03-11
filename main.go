package main

import (
	"fmt"
	"golearning/algorithms"
	"golearning/basegrammar"
)

func main() {
	//变量测试
	//basegrammar.VariableAndConst()
	//字符串遍历
	//basegrammar.RandString()
	//数组测试
	//basegrammar.ArrTest()
	//切片测试
	//basegrammar.SliceTest()
	//测试map
	fmt.Println("测试map")
	basegrammar.MapTest()
	//测试结构体
	//basegrammar.StruceTest()
	//basegrammar.StruceTest()
	//测试流程控制
	//basegrammar.FlowTest()
	//测试函数
	//basegrammar.TestFun()
	//basegrammar.TestInterface()

	//####################算法#######################//
	//
	//basesort.TestSort()
	//常见库
	//baselib.TestLib()
	//basepkg.ReadByViper()
	//接口
	//baseobject.TestObject()
	//协程
	//basegoroutine.Run()
	//树
	//tree.Test
	//链表
	//linkedlist.Test()
	//	栈
	//stack.Test()
	//队列
	//queue.Test()
	//排序
	fmt.Println("\n测试排序")
	algorithms.Test()
}
