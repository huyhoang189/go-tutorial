package Observer

import "fmt"

type Item struct {
	observers []Observer
	Name      string
	InStock   bool
}

func NewItem(name string) *Item {
	return &Item{Name: name}
}

func (i *Item) UpdateAvailability() {
	fmt.Printf("Item %s is now in stock \n ", i.Name)
	i.InStock = true
	i.NotifyAll()
}

func (i *Item) NotifyAll() {
	for _, o := range i.observers {
		o.Update(i.Name)
	}
}

func (i *Item) Register(o Observer) {
	i.observers = append(i.observers, o)
}

func (i *Item) DeRegister(o Observer) {
	i.observers = removeObserver(o, i.observers)
}

func removeObserver(o Observer, os []Observer) []Observer {
	l := len(os)
	for i, e := range os {
		if e.GetID() == o.GetID() {
			os[l-1], os[i] = os[i], os[l-1]
			return os[:l-1]
		}
	}
	return os
}
