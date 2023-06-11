package Visitor

type Shape interface {
	Accept(Visitor) float32
}
