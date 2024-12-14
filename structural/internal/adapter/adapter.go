package adapter

import "fmt"

// Computer interface
type Computer interface {
	InsertIntoLightningPort()
}

// Implementation of Client
type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("> Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

// Implementation of Target - Mac
type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

// Implementation of Adaptee - Windows
type Windows struct {
}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

// Adapter from Mac to Windows
type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}

// Factory method for create Windows Adapter
func NewWindowsAdapter(windowMachine *Windows) WindowsAdapter {
	return WindowsAdapter{windowMachine: windowMachine}
}
