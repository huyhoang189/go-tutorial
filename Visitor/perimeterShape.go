package Visitor

type Perimeter struct {
}

func (per *Perimeter) VisitForSquare(s *Square) float32 {
	return float32(s.Side * 4)
}
func (per *Perimeter) VisitForCircle(c *Circle) float32 {
	return float32(c.Radius) * 2 * PI
}
