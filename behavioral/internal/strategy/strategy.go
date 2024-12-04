package strategy

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

// Strategy interface
type RouteStrategy interface {
	BuildRoute(A Point, B Point)
}

// Implementation of the route by Taxi
type TaxiRoute struct {
}

func (s *TaxiRoute) BuildRoute(A Point, B Point) {
	fmt.Printf("Taxi route from A=(%d; %d) to B=(%d; %d)\n", A.X, A.Y, B.X, B.Y)
}

// Implementation of the route by Bus
type BusRoute struct {
}

func (s *BusRoute) BuildRoute(A Point, B Point) {
	fmt.Printf("Bus route from A=(%d; %d) to B=(%d; %d)\n", A.X, A.Y, B.X, B.Y)
}

// Implementation of the route by Bike
type BikeRoute struct {
}

func (s *BikeRoute) BuildRoute(A Point, B Point) {
	fmt.Printf("Bike route from A=(%d; %d) to B=(%d; %d)\n", A.X, A.Y, B.X, B.Y)
}

// Context for choose RouteStrategy and build route
type RouteCreator struct {
	routeStrategy RouteStrategy
}

func NewRouteCreator(routeStrategy RouteStrategy) RouteCreator {
	return RouteCreator{routeStrategy: routeStrategy}
}

func (r *RouteCreator) SetStrategy(strategy RouteStrategy) {
	r.routeStrategy = strategy
}

func (r *RouteCreator) BuildRoute(A Point, B Point) {
	r.routeStrategy.BuildRoute(A, B)
}
