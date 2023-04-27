package baseobject

/**
方法(类)
*/

type Student struct {
	id    uint
	name  string
	score float64
}

func NewStudent(id uint, name string, score float64) *Student {
	return &Student{id: id, name: name, score: score}
}
func NewStudentV2(id uint, name string, score float64) Student {
	return Student{id: id, name: name, score: score}
}
func (s Student) GetName() string {
	return s.name
}
func (s *Student) SetName(name string) {
	s.name = name
}
