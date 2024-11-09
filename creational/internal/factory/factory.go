package factory

import "errors"

// Transport interface, struct and his methods
type TransportInterface interface {
	GetName() string
	GetSpeed() float64
	SetName(newName string)
	SetSpeed(newSpeed float64)
	GetNumberOfEngines() int
	GetAxleLoad() int
}

type transport struct {
	name  string
	speed float64
}

func (t *transport) GetName() string {
	return t.name
}

func (t *transport) GetSpeed() float64 {
	return t.speed
}

func (t *transport) SetName(newName string) {
	t.name = newName
}

func (t *transport) SetSpeed(newSpeed float64) {
	t.speed = newSpeed
}

func (t *transport) GetNumberOfEngines() int {
	return -1
}

func (t *transport) GetAxleLoad() int {
	return -1
}

// Ship struct (with inline field) and his methods
type Ship struct {
	transport
}

func (t *Ship) GetNumberOfEngines() int {
	return 4
}

// Truck struct (with inline field) and his methods
type Truck struct {
	transport
}

func (t *Truck) GetAxleLoad() int {
	return 10
}

// Create methods for Ship
func newShip() TransportInterface {
	return &Ship{
		transport: transport{
			name:  "Sea Ship",
			speed: 40.5,
		},
	}
}

// Create methods for Truck
func newTruck() TransportInterface {
	return &Truck{
		transport: transport{
			name:  "Road Truck",
			speed: 71.7,
		},
	}
}

// Factory method
func GetTransport(transportType string) (TransportInterface, error) {
	if transportType == "ship" {
		return newShip(), nil
	} else if transportType == "truck" {
		return newTruck(), nil
	}
	return nil, errors.New("wrong transport type passed")
}
