package main

import (
	"design-patterns/creational/internal/absfactory"
	"fmt"
)

func printCrossoverDetails(crossover absfactory.CrossoverInterface) {
	fmt.Printf("Crossover - brand: %s, speed: %f\n", crossover.GetBrand(), crossover.GetSpeed())
}

func printSedanDetails(sedan absfactory.SedanInterface) {
	fmt.Printf("Sedan - brand: %s, speed: %f\n", sedan.GetBrand(), sedan.GetSpeed())
}

func main() {
	volkswagenFactory, err := absfactory.GetCarsFactory("volkswagen")
	if err != nil {
		panic(err.Error())
	}
	volkswagenSedan := volkswagenFactory.CreateSedan()
	volkswagenCrossover := volkswagenFactory.CreateCrossover()
	fmt.Println("<VolkswagenFactory>")
	printSedanDetails(volkswagenSedan)
	printCrossoverDetails(volkswagenCrossover)

	toyotaFactory, err := absfactory.GetCarsFactory("toyota")
	if err != nil {
		panic(err.Error())
	}
	toyotaSedan := toyotaFactory.CreateSedan()
	toyotaCrossover := toyotaFactory.CreateCrossover()
	fmt.Println("<ToyotaFactory>")
	printSedanDetails(toyotaSedan)
	printCrossoverDetails(toyotaCrossover)
}
