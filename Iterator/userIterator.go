package Iterator

type UserIterator struct {
	users []*User
	index int
}

func (ui *UserIterator) HasNext() bool {
	if ui.index < len(ui.users) {
		return true
	}
	return false
}

func (ui *UserIterator) GetNext() *User {
	if ui.HasNext() {
		user := ui.users[ui.index]
		ui.index++
		return user
	}
	return nil
}
