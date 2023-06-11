package Visitor

type Circle struct {
	Radius int
}

func (c *Circle) Accept(v Visitor) float32 {
	return v.VisitForCircle(c)
}
