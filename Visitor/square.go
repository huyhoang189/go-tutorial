package Visitor

type Square struct {
	Side int
}

func (s *Square) Accept(v Visitor) float32 {
	return v.VisitForSquare(s)
}
