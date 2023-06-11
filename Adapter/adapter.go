package Adapter

type MacAdapter struct {
	M Mac
}

func (ma *MacAdapter) insertByUsbA() {
	ma.M.insertByUsbC()
}
