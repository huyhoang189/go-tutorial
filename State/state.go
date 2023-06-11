package State

type State interface {
	On(m *Machine)
	Off(m *Machine)
}
