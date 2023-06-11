package Mementor

type Memento struct {
	State string
}

func (m *Memento) GetSaveState() string {
	return m.State
}
