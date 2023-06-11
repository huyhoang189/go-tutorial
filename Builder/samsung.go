package Builder

type Samsung struct {
	Phone Phone
}

func (s *Samsung) GetPhone() Phone {
	return s.Phone
}

func (s *Samsung) SetName() BuildProcess {
	s.Phone.Name = "Samsung"
	return s
}
func (s *Samsung) SetCamera() BuildProcess {
	s.Phone.Camera = true
	return s
}

func (s *Samsung) SetTorch() BuildProcess {
	s.Phone.Torch = false
	return s
}

func (s *Samsung) SetColorDisplay() BuildProcess {
	s.Phone.ColorDisplay = true
	return s
}

func (s *Samsung) SetDualSim() BuildProcess {
	s.Phone.DualSim = false
	return s
}
