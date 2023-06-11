package Visitor

type Visitor interface {
	VisitForSquare(s *Square) float32
	VisitForCircle(c *Circle) float32
}
