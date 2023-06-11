package Bridge

import "fmt"

type HP struct {
}

func (e *HP) printing(str string) {
	fmt.Println("HP : printing document from ", str)
}
