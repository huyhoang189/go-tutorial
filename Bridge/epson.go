package Bridge

import "fmt"

type EPSON struct {
}

func (e *EPSON) printing(str string) {
	fmt.Println("EPSON : printing document from ", str)
}
