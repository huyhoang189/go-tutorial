package Prototype

type INode interface {
	Clone() INode
	Print(s string)
}
