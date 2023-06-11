package Mediator

import "fmt"

type PassengerTrain struct {
	Mediator Mediator
}

func (pt *PassengerTrain) RequestArrival() {
	if pt.Mediator.CanLand(pt) {
		fmt.Println("Passenger Train : Lading")
	} else {
		fmt.Println("Passenger Train : Waiting")
	}

}

func (pt *PassengerTrain) Departure() {
	fmt.Println("Passenger Train : Living")
	pt.Mediator.NotifyFree()
}

func (pt *PassengerTrain) PermitArrival() {
	fmt.Println("Passenger Train : Arrival Permitted. Lading")
}
