package Iterator

type Collection interface {
	createIterator() Iterator
}
