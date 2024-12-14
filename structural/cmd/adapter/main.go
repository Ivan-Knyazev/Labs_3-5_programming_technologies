package main

import (
	"design-patterns/structural/internal/adapter"
)

func main() {

	client := adapter.Client{}
	mac := adapter.Mac{}

	client.InsertLightningConnectorIntoComputer(&mac)

	windowsMachine := adapter.Windows{}
	windowsMachineAdapter := adapter.NewWindowsAdapter(&windowsMachine)

	client.InsertLightningConnectorIntoComputer(&windowsMachineAdapter)
}
