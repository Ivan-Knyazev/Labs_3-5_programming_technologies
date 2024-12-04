package main

import "design-patterns/behavioral/internal/strategy"

func main() {
	taxiStrategy := strategy.TaxiRoute{}
	creator := strategy.NewRouteCreator(&taxiStrategy)

	// Create route with taxiStrategy
	creator.BuildRoute(strategy.Point{X: 2, Y: 3}, strategy.Point{X: 30, Y: 41})

	// Create route with busStrategy
	busStrategy := strategy.BusRoute{}
	creator.SetStrategy(&busStrategy)
	creator.BuildRoute(strategy.Point{X: 1, Y: 7}, strategy.Point{X: 20, Y: 31})

	// Create route with bikeStrategy
	bikeStrategy := strategy.BikeRoute{}
	creator.SetStrategy(&bikeStrategy)
	creator.BuildRoute(strategy.Point{X: 12, Y: 9}, strategy.Point{X: 19, Y: 17})
}
