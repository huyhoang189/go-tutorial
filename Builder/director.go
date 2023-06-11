package Builder

type Director struct {
	builder BuildProcess
}

func (d *Director) SetBuilder(b BuildProcess) {
	d.builder = b
}
func (d *Director) BuildPhone() Phone {
	d.builder.SetCamera().SetDualSim().SetTorch().SetColorDisplay().SetName()
	return d.builder.GetPhone()
}
