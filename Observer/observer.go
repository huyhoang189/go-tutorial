package Observer

type Observer interface {
	GetID() string
	Update(str string)
}
