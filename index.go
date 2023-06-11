package main

import (
	"fmt"
	"main/AbstractFactory"
	"main/Adapter"
	"main/Bridge"
	"main/Builder"
	"main/Composite"
	"main/Mediator"
	"main/Mementor"
	"main/Observer"
	"main/Prototype"
	"main/Singleton"
	"main/State"
	"main/Strategies"
	"main/TemplateMethod"
	"main/Visitor"
)

func main() {
	/*Design Pattern
	1. Singleton : cpSingleton()
	2. Builder : cpBuilder()
	3. Abstract Factory : cpAbstractFactory()
	4. Prototype : cpPrototype()
	5. Iterator : bpIterator()
	6. Mediator : bpMediator()
	7. Memento : bpMemento()
	8. Observer : bpObserver()
	9. State : bpState()
	10. Strategy : bpStrategy()
	11. Template method : bpTemplateMethod()
	12. Visitor : bpVisitor()
	13. Adapter : spAdapter()
	13. Bridge : spBridge()
	14. Composite : spComposite()
	15. Flyweight :  spFlyweight()
	*/
	//bpStrategy()

	//bpVisitor()
	spFlyweight()

}

// for abstract factory
func fmtShoe(shoe AbstractFactory.IShoe) {
	fmt.Printf("%v %v %v \n", shoe.GetName(), shoe.GetSize(), shoe.GetBrand())
}

func fmtShort(short AbstractFactory.IShort) {

	fmt.Printf("%v %v %v \n", short.GetName(), short.GetSize(), short.GetBrand())
}

// for memento pattern
func addItem(ct *Mementor.CareTaker, origin *Mementor.Originator, s string) {
	origin.SetState(s)
	ct.AddMemento(origin.CreateMemento())
}

func fmtItem(origin *Mementor.Originator) {
	fmt.Println("Current state is : ", origin.GetState())
}

// BP_Observer
func bpObserver() {
	shirtItem := Observer.Item{Name: "Nike shirt"}
	observerFirst := &Observer.Customer{ID: "huyhoang189@gmail.com"}
	observerSecond := &Observer.Customer{ID: "vitcon1889@gmail.com"}

	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)
	shirtItem.DeRegister(observerFirst)
	shirtItem.UpdateAvailability()
}

func bpMemento() {
	careTaker := &Mementor.CareTaker{
		Mementos: make([]*Mementor.Memento, 0),
	}
	originator := &Mementor.Originator{}

	addItem(careTaker, originator, "A")
	fmtItem(originator)
	addItem(careTaker, originator, "B")
	fmtItem(originator)
	addItem(careTaker, originator, "C")
	fmtItem(originator)
	//fmt.Println("", careTaker.Mementos)

	fmt.Println("Back item")
	originator.Restore(careTaker.GetMemento(1))
	fmtItem(originator)
	originator.Restore(careTaker.GetMemento(0))
	fmtItem(originator)
}

func bpMediator() {
	stationManager := Mediator.NewStationManager()
	passengerTrain := &Mediator.PassengerTrain{
		Mediator: stationManager,
	}
	goodTrain := &Mediator.GoodTrain{
		Mediator: stationManager,
	}

	passengerTrain.RequestArrival()
	goodTrain.RequestArrival()
	passengerTrain.Departure()
}

func cpPrototype() {
	file1 := &Prototype.File{Name: "file 1"}
	file2 := &Prototype.File{Name: "file 2"}
	file3 := &Prototype.File{Name: "file 3"}
	folder1 := &Prototype.Folder{Name: "folder 1", Childrens: []Prototype.INode{file1}}

	folder2 := &Prototype.Folder{Name: "folder 1", Childrens: []Prototype.INode{folder1, file2, file3}}

	folder2.Print("   ")
	folder3 := folder2.Clone()
	folder3.Print("   ")
}

func cpAbstractFactory() {
	factory := AbstractFactory.GetSportsFactory("adidas")
	adidas := factory.MakeShoe()
	fmtShoe(adidas)

	factory = AbstractFactory.GetSportsFactory("nike")
	nike := factory.MakeShoe()
	fmtShoe(nike)
}

func cpBuilder() {
	direct := Builder.Director{}
	iphone := &Builder.Iphone{}
	direct.SetBuilder(iphone)
	phone := direct.BuildPhone()

	fmt.Println(phone)

	samsung := &Builder.Samsung{}
	direct.SetBuilder(samsung)
	phone = direct.BuildPhone()

	fmt.Println(phone)
}

func cpSingleton() {
	var s = Singleton.GetInstance()
	var s1 = Singleton.GetInstance()
	fmt.Printf("%p\n", s)
	fmt.Printf("%p\n", s1)
}

func bpState() {
	fmt.Println("Pattern State")
	machine := State.NewMachine()
	machine.OnMachine()
	machine.OffMachine()
}

func bpStrategy() {
	fmt.Println("Pattern strategy")
	lfu := &Strategies.Lfu{}
	fifo := &Strategies.Fifo{}
	lru := &Strategies.Lfu{}

	cache := Strategies.InitCache(lfu)

	cache.Add("1", "A")
	cache.Add("2", "B")
	cache.Add("3", "C")
	cache.SetEvictionAlgo(fifo)
	cache.Add("4", "D")
	cache.SetEvictionAlgo(lru)
	cache.Add("5", "E")

}

func bpTemplateMethod() {
	sms := &TemplateMethod.SMS{}
	email := &TemplateMethod.Email{}

	otp := &TemplateMethod.Otp{
		ObjOTP: sms,
	}

	err := (*otp).GenAndSendOTP(4)
	if err != nil {
		return
	}

	otp = &TemplateMethod.Otp{
		ObjOTP: email,
	}

	err = otp.GenAndSendOTP(4)
	if err != nil {
		return
	}
}

func bpVisitor() {
	circle := Visitor.Circle{Radius: 5}
	square := Visitor.Square{Side: 10}

	acreage := &Visitor.Acreage{}
	perimeter := &Visitor.Perimeter{}

	fmt.Println(circle.Accept(acreage))
	fmt.Println(circle.Accept(perimeter))

	fmt.Println(square.Accept(acreage))
	fmt.Println(square.Accept(perimeter))
}

func spAdapter() {
	computer := &Adapter.MacAdapter{}

	client := Adapter.Client{}
	client.InsertPortToComputer(computer)
}

func spBridge() {
	printerEpson := &Bridge.EPSON{}
	printerHp := &Bridge.HP{}

	win := Bridge.Windows{}

	win.SetPrinter(printerEpson)
	win.PrintDocument()

	win.SetPrinter(printerHp)
	win.PrintDocument()

	mac := Bridge.Macos{}

	mac.SetPrinter(printerEpson)
	mac.PrintDocument()

	mac.SetPrinter(printerHp)
	mac.PrintDocument()
}

func spComposite() {
	file1 := &Composite.File{Name: "file 1"}
	file2 := &Composite.File{Name: "file 2"}
	file3 := &Composite.File{Name: "file 3"}
	file4 := &Composite.File{Name: "file 4"}
	//file5 := &Composite.File{Name: "file 5"}
	//file6 := &Composite.File{Name: "file 6"}
	//file7 := &Composite.File{Name: "file 7"}

	folder1 := &Composite.Folder{Name: "folder 1"}
	folder1.AddComponent(file1)
	folder1.AddComponent(file2)

	folder2 := &Composite.Folder{Name: "folder 2"}
	folder2.AddComponent(folder1)
	folder2.AddComponent(file3)
	folder2.AddComponent(file4)

	folder2.Search("file 1")

}

func spFlyweight() {
	
}
