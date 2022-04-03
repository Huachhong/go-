package basegrammar

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	name, city string
	age uint
}

func NewPerson(name,city string, age uint) *Person {
	return &Person{
		city: city,
		name: name,
		age: age,
	}
}

//方法与接收者(接收者可以是值类型，也可以是指针类型)

func (p *Person) SetAge(age uint) {
	p.age = age
}

//嵌套结构体(继承
type liveb struct {
	foot uint8
}
func (b *liveb) sayFoots() {
	fmt.Printf("%d条腿\n", b.foot)
}
type Animal struct {
	name string
	*liveb
}
func (a *Animal) move() {
	fmt.Printf("%s会动 \n", a.name)
}
type Dog struct {
	feet uint8
	*Animal
}
func (d *Dog) wang() {
	fmt.Printf("%s会汪汪", d.Animal.name)
}

func TestInherit() {
	//p1 := NewPerson("萧红", "北京", 28)
	p1 := &Person{"小哈", "西安", 15}
	p1.SetAge(30)
	fmt.Printf("\n set age p1 = %v", p1)
	fmt.Println("继承")
	d := &Dog{
		feet:2,
		Animal: &Animal{"小狗", &liveb{4}},
	}
	d.sayFoots()
	d.move()
	d.wang()
}

//结构体与json序列化
type Student struct {
	ID int
	Gender, Name string
}
type Class struct {
	Title string
	Students []*Student
}
func StructJson() {
	c1 := &Class{
		Title: "101",
		Students: make([]*Student, 0 , 10),
	}
	for i := 0; i < 10; i++  {
		s := &Student{
			ID: i,
			Name: fmt.Sprintf("stu%02d", i),
			Gender: "男",
		}
		c1.Students = append(c1.Students, s)
	}
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Printf("转换json有误, err=%s", err)
	}
	fmt.Printf("\n class json = %s", data)
}

//结构体与标签
type Teacher struct {
	ID int `json:"id"` //通过指定tag实现json序列化该字段时的key
	Name string
	Gender string
	age int //私有不能被json包访问
}

func StructTagInJson() {
	ts := make([]Teacher, 0, 10)
	for i := 0; i < 10; i++   {
		t := Teacher{
			ID: i,
			Name: fmt.Sprintf("Te%02d", i),
			Gender: "男",
			age: i * 2, //
		}
		ts = append(ts, t)
	}
	data, _ := json.Marshal(ts)
	fmt.Printf("\n struct tag to json =%s", data)
}

func StruceTest() {
	fmt.Printf("\n测试结构体")
	var p1 Person //结构体实例化
	p1.name ="萧红"
	p1.city = "深圳"
	p1.age = 20
	fmt.Printf("\np1 = %v", p1)

	var u struct{name string; addr string} //匿名结构体
	u.name = "Mr Li"
	u.addr = "广州"
	fmt.Printf("\nu=%v", u)
	//new实例化
	var p2 = new(Person) //返回指针
	p2.name = "hong"
	p2.age  = 25
	p2.city = "深圳"
	fmt.Printf("\nu=%v, u prt = %T", p2, p2)

	//取结构体的地址实例化 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作
	var p3 = &Person{}
	p3.city = "上海"
	p3.age = 100
	//p3.name = "滨海路"
	fmt.Printf("\np3=%#v", p3)

	p4 := NewPerson("小希", "上海", 100)
	fmt.Printf("\np4=%#v", p4)

	fmt.Printf("\n测试结构体和json")
	StructJson()
	StructTagInJson()
}