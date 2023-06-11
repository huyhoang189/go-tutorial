package State

import "fmt"

type Machine struct {
	Current State
}

func NewMachine() *Machine {
	fmt.Println("Machine is ready.")
	return &Machine{
		NewOff(),
	}
}

func (m *Machine) SetCurrent(s State) {
	m.Current = s
}

func (m *Machine) OnMachine() {
	m.Current.On(m)
}

func (m *Machine) OffMachine() {
	m.Current.Off(m)
}
