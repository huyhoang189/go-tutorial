package Adapter

import "fmt"

type Mac struct {
}

func (m *Mac) insertByUsbC() {
	fmt.Println("Insert Mac by usb c")
}
