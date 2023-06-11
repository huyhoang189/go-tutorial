package Mediator

type Train interface {
	RequestArrival()
	Departure()
	PermitArrival()
}
