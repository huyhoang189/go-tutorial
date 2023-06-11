package AbstractFactory

// IShoe is an interface that defines methods for setting and getting shoe attributes
type IShoe interface {
	setName(str string)
	setSize(i int)
	setBrand(str string)
	GetName() string
	GetSize() int
	GetBrand() string
}

// shoe is a struct that represents a shoe object with name, size, and brand attributes
type shoe struct {
	name  string
	size  int
	brand string
}

// setName is a receiver function that sets the name attribute of a shoe object
func (s *shoe) setName(str string) {
	s.name = str
}

// setSize is a receiver function that sets the size attribute of a shoe object
func (s *shoe) setSize(i int) {
	s.size = i
}

// setBrand is a receiver function that sets the brand attribute of a shoe object
func (s *shoe) setBrand(str string) {
	s.brand = str
}

// GetName is a receiver function that returns the name attribute of a shoe object
func (s *shoe) GetName() string {
	return s.name
}

// GetSize is a receiver function that returns the size attribute of a shoe object
func (s *shoe) GetSize() int {
	return s.size
}

// GetBrand is a receiver function that returns the brand attribute of a shoe object
func (s *shoe) GetBrand() string {
	return s.brand
}
