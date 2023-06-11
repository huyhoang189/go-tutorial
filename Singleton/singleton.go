package Singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var (
	instance *singleton
)

func init() {
	instance = &singleton{count: 100}
}

func GetInstance() Singleton {
	return instance
}

func (s *singleton) AddOne() int {
	s.count = s.count + 1
	return s.count
}
