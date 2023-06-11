package Builder

type BuildProcess interface {
	SetName() BuildProcess
	SetCamera() BuildProcess
	SetDualSim() BuildProcess
	SetTorch() BuildProcess
	SetColorDisplay() BuildProcess
	GetPhone() Phone
}
