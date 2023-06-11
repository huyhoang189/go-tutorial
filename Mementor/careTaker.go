package Mementor

type CareTaker struct {
	Mementos []*Memento
}

func (c *CareTaker) AddMemento(m *Memento) {
	c.Mementos = append(c.Mementos, m)
}

func (c *CareTaker) GetMemento(index int) *Memento {
	return c.Mementos[index]
}
