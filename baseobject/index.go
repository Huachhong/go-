package baseobject

import "fmt"

func TestObject() {
	var m Mover
	d := &Dog{}
	m = d
	m.move()

	s := NewStudent(1, "LiMing", 100)
	//正常调用指针方法
	s.SetName("芳芳")
	fmt.Println(s.GetName())

	s1 := NewStudentV2(2, "LiSi", 99)
	s1.SetName("Hk Name")
	fmt.Println(s1.GetName())

	NewStudent(1, "HH", 80).SetName("Kong")

}
