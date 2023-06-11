package Bridge

type Macos struct {
	printer Printer
}

func (w *Macos) PrintDocument() {
	w.printer.printing("Macos")
}

func (w *Macos) SetPrinter(printer Printer) {
	w.printer = printer
}
