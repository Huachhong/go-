package baselib

import (
	"fmt"
	"golearning/basegrammar"
	"os"
)

//常见标准库及使用

//Print系列
func testPrint() {
	fmt.Println("##########Print系列")
	fmt.Print("我是世界")
	name := "古藤老树昏鸦"
	fmt.Printf("name = %s\n", name)
	//fmt.Println(name)
	fmt.Println("在终端打印单独一行显示")
	line := basegrammar.Tiger{"老虎"}
	fmt.Printf("%#v", line)
}
func testFprint() {
	fmt.Println("##########Fprint系列")
	fileObj, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错")
	}
	name := "床前明月光"
	fmt.Fprintf(fileObj, "##%s\n", name)
}

func testSprint() {
	fmt.Println("##########Sprint系列")
	str := "静夜系数"
	fs := fmt.Sprintln(str)
	fmt.Print(fs)
}

func testErrorF() {
	fmt.Println("##########Errorf")
	st := "this is a error"
	err := fmt.Errorf("error:%s", st)
	fmt.Println(err)
}

func testZan() {
	fmt.Println("##########测试格式占位符")
	f := 88.88
	fmt.Printf("%f\n", f)
	fmt.Printf("%10f\n", f)
	fmt.Printf("%.2f\n", f)
	fmt.Printf("%10.2f\n", f)
	fmt.Printf("%10.f\n", f)
}

func testStdIn() {
	fmt.Println("##########标准输入测试")
	var (
		name    string
		age     int
		married bool
	)
	//fmt.Scan(&name, &age, &married)
	fmt.Printf("name=%s, age=%d, married=%t\n", name, age, married)

	var (
		head string
		foot int
		hand int
	)
	fmt.Scanf("h:%s,f:%d,h:%d", &head, &foot, &hand)
	fmt.Printf("head:%s, head:%d, hand:%d", head, foot, hand)
}

func TestSPrint() {
	testPrint()
	testFprint()
	testSprint()
	testErrorF()
	testZan()
	testStdIn()
}
