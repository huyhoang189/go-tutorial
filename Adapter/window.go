package Adapter

import "fmt"

type Window struct {
}

func (w *Window) insertByUsbA() {
	fmt.Println("Insert window by usb A")
}
