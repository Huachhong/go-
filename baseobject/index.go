package baseobject

func TestObject() {
	var m Mover
	d := &Dog{}
	m = d
	m.move()
}
