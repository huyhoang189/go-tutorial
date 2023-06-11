package Mediator

import "fmt"

type GoodTrain struct {
	Mediator Mediator
}

func (gt *GoodTrain) RequestArrival() {
	if gt.Mediator.CanLand(gt) {
		fmt.Println("Good Train : Lading")
	} else {
		fmt.Println("Good Train : Waiting")
	}

}

func (gt *GoodTrain) Departure() {
	fmt.Println("Good Train : Living")
	gt.Mediator.NotifyFree()
}

func (gt *GoodTrain) PermitArrival() {
	fmt.Println("Good Train : Arrival Permitted. Lading")
}
