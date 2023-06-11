package Bridge

type Computer interface {
	PrintDocument()
	SetPrinter(Printer)
}
