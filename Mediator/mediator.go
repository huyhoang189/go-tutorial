package Mediator

type Mediator interface {
	CanLand(Train) bool
	NotifyFree()
}
