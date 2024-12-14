package bridge

import (
	"fmt"
)

// Implementor
type Printer interface {
	PrintFile()
}

// Concrete Implementor - Epson
type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

// Concrete Implementor - HP
type HP struct {
}

func (p *HP) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}

// Abstraction
type Computer interface {
	Print()
	SetPrinter(Printer)
}

// Refined Abstraction - Mac
type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

// Refined Abstraction - Windows
type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}
