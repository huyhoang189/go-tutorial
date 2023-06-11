package State

import "fmt"

type Off struct {
}

func NewOff() State {
	return &Off{}
}

func (o *Off) Off(m *Machine) {
	fmt.Println("Machine is already off")
}

func (o *Off) On(m *Machine) {
	fmt.Println("Machine is going Off to On")
	m.SetCurrent(NewOn())
}
