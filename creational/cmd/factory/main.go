package main

import (
	"design-patterns/creational/internal/factory"
	"fmt"
)

func printDetails(transport factory.TransportInterface) {
	fmt.Printf("Transport: %s, speed: %f, ", transport.GetName(), transport.GetSpeed())
	switch transport.(type) {
	case *factory.Ship:
		fmt.Printf("*number of engines: %d\n", transport.GetNumberOfEngines())
	case *factory.Truck:
		fmt.Printf("*axle load: %d\n", transport.GetAxleLoad())
	}
}

func main() {
	ship, err := factory.GetTransport("ship")
	if err != nil {
		panic(err.Error())
	}
	printDetails(ship)
	truck, err := factory.GetTransport("truck")
	if err != nil {
		panic(err.Error())
	}
	printDetails(truck)
}
