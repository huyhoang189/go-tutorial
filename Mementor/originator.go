package Mementor

type Originator struct {
	State string
}

func (o *Originator) CreateMemento() *Memento {
	return &Memento{
		State: o.State,
	}
}

func (o *Originator) Restore(m *Memento) {
	o.State = m.State
}

func (o *Originator) GetState() string {
	return o.State
}

func (o *Originator) SetState(s string) {
	o.State = s
}
