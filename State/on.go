package State

import "fmt"

type On struct {
}

func NewOn() State {
	return &On{}
}

func (o *On) On(m *Machine) {
	fmt.Println("Machine already on")

}

func (o *On) Off(m *Machine) {
	fmt.Println("Machine is going from On to Off")
	m.SetCurrent(NewOff())
}
