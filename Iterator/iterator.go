package Iterator

type Iterator interface {
	HasNext() bool
	GetNext() *User
}
