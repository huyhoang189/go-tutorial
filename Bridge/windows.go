package Bridge

type Windows struct {
	printer Printer
}

func (w *Windows) PrintDocument() {
	w.printer.printing("windows")
}

func (w *Windows) SetPrinter(printer Printer) {
	w.printer = printer
}
