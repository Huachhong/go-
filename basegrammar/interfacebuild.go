package basegrammar

import "fmt"

type Sayer interface {
	say() string
}

type Tiger struct {
	Name string
}

//值接收者实现接口
func (t Tiger) say() (string) {
	fmt.Printf("\nsay = %v", t.Name)
	return t.Name
}

type Duck struct {
Name string
}
//指针接收者实现接口
func (d *Duck) say() string {
	fmt.Printf("\nsay = %v", d.Name)
	return  d.Name
}

func TestEmptyInterface(a interface{}) {
	fmt.Printf("\ntype = %T, value = %v", a, a)
}

func TestInterface() {
	t := Tiger{"老虎"} //t 是Tiger类型
	var sa Sayer
	sa = t //sa 可以接收Tiger类型
	sa.say()
	ch := &Tiger{"老虎幼崽"} //ch 是*Tiger类型
	var sch Sayer
	sch = ch //sch可以接收*Tiger类型
	sch.say()
	//以上是测试值接收者实现接口

	d := &Duck{"鸭子"} //d 是*Duck类型
	var dch Sayer
	dch = d //dch只能接收*Duck类型，不能接收Duck类型
	dch.say()

	//空接口
	TestEmptyInterface(Tiger{"白色的老虎"})
	//var ai map[string]interface{}
	ai := make(map[string]interface{}, 1)
	ai["name"] = "老虎"
	ai["color"] = "white"
	ai["age"] = 5
	ai["status"] = false
	TestEmptyInterface(ai)
	//断言
	var f interface{}
	f = "pprod"
	re, ok := f.(int)
	fmt.Printf("\nresult=%v, ok=%v", re, ok)

}