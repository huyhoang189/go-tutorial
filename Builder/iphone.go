package Builder

type Iphone struct {
	Phone Phone
}

func (ip *Iphone) SetCamera() BuildProcess {
	ip.Phone.Camera = false
	return ip
}

func (ip *Iphone) SetDualSim() BuildProcess {
	ip.Phone.DualSim = false
	return ip
}

func (ip *Iphone) SetTorch() BuildProcess {
	ip.Phone.Torch = true
	return ip
}

func (ip *Iphone) SetColorDisplay() BuildProcess {
	ip.Phone.ColorDisplay = false
	return ip
}
func (ip *Iphone) SetName() BuildProcess {
	ip.Phone.Name = "Iphone"
	return ip
}

func (ip *Iphone) GetPhone() Phone {
	return ip.Phone
}
