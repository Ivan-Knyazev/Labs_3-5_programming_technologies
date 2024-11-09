package absfactory

import "errors"

// Abstract types of cars with his methods
type SedanInterface interface {
	GetBrand() string
	GetSpeed() float64
	SetSpeed(newSpeed float64)
}

type sedan struct {
	brand string
	speed float64
}

func (s *sedan) GetBrand() string {
	return s.brand
}

func (s *sedan) GetSpeed() float64 {
	return s.speed
}

func (s *sedan) SetSpeed(newSpeed float64) {
	s.speed = newSpeed
}

type CrossoverInterface interface {
	GetBrand() string
	GetSpeed() float64
	SetSpeed(newSpeed float64)
}

type crossover struct {
	brand string
	speed float64
}

func (s *crossover) GetBrand() string {
	return s.brand
}

func (s *crossover) GetSpeed() float64 {
	return s.speed
}

func (s *crossover) SetSpeed(newSpeed float64) {
	s.speed = newSpeed
}

// Specific cars of defferent types
type VolkswagenSedan struct {
	sedan
}

type ToyotaSedan struct {
	sedan
}

type VolkswagenCrossover struct {
	crossover
}

type ToyotaCrossover struct {
	crossover
}

// Abstract Factory
type CarsFactoryInterface interface {
	CreateSedan() SedanInterface
	CreateCrossover() CrossoverInterface
}

// Specific cars factories and his methods of defferent types
type VolkswagenFactory struct {
}

func (s *VolkswagenFactory) CreateSedan() SedanInterface {
	return &VolkswagenSedan{
		sedan: sedan{
			brand: "Volkswagen",
			speed: 162.7,
		},
	}
}

func (s *VolkswagenFactory) CreateCrossover() CrossoverInterface {
	return &VolkswagenCrossover{
		crossover: crossover{
			brand: "Volkswagen",
			speed: 182.7,
		},
	}
}

type ToyotaFactory struct {
}

func (s *ToyotaFactory) CreateSedan() SedanInterface {
	return &ToyotaSedan{
		sedan: sedan{
			brand: "Toyota",
			speed: 210.3,
		},
	}
}

func (s *ToyotaFactory) CreateCrossover() CrossoverInterface {
	return &ToyotaCrossover{
		crossover: crossover{
			brand: "Toyota",
			speed: 230.3,
		},
	}
}

// Function for create Factory
func GetCarsFactory(brand string) (CarsFactoryInterface, error) {
	if brand == "volkswagen" {
		return &VolkswagenFactory{}, nil
	} else if brand == "toyota" {
		return &ToyotaFactory{}, nil
	}
	return nil, errors.New("wrong transport type passed")
}
