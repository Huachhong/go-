package baseobject

import "fmt"

type Dog struct {
}

func (d *Dog) move() {
	fmt.Println("狗狗移动")
}
