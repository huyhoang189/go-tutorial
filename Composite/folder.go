package Composite

import "fmt"

type Folder struct {
	Name       string
	components []Component
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Searching keyword %s in folder %s \n", keyword, f.Name)
	for _, component := range f.components {
		component.Search(keyword)
	}
}

func (f *Folder) AddComponent(c Component) {
	f.components = append(f.components, c)
	//fmt.Println("Add")
}
