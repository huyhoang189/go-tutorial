package Prototype

import "fmt"

type Folder struct {
	Name      string
	Childrens []INode
}

func (f *Folder) Print(s string) {
	fmt.Println(s + f.Name)
	for _, e := range f.Childrens {
		e.Print(s + s)
	}
}

func (f *Folder) Clone() INode {
	cloneFolder := &Folder{
		Name: f.Name + "_Clone",
	}

	var tempChildrens []INode

	for _, e := range f.Childrens {
		copy := e.Clone()
		tempChildrens = append(tempChildrens, copy)
	}

	cloneFolder.Childrens = tempChildrens
	return cloneFolder
}
