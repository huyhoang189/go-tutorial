package AbstractFactory

type IShort interface {
	setName(str string)
	setSize(i int)
	setBrand(str string)
	GetName() string
	GetSize() int
	GetBrand() string
}

type short struct {
	name  string
	size  int
	brand string
}

func (s *short) setName(str string) {
	s.name = str
}
func (s *short) setSize(i int) {
	s.size = i
}

func (s *short) setBrand(str string) {
	s.brand = str
}

func (s *short) GetName() string {
	return s.name
}
func (s *short) GetSize() int {
	return s.size
}
func (s *short) GetBrand() string {
	return s.brand
}
