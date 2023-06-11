package Visitor

const PI = 3.14

type Acreage struct {
}

func (acr *Acreage) VisitForSquare(s *Square) float32 {
	return float32(s.Side * s.Side)
}
func (acr *Acreage) VisitForCircle(c *Circle) float32 {
	return float32(c.Radius) * float32(c.Radius) * PI
}
