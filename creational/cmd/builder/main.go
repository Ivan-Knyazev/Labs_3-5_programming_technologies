package main

import "design-patterns/creational/internal/builder"

func main() {
	defaultBuilder, err := builder.NewBuilder("default")
	if err != nil {
		panic(err)
	}
	gamingBuilder, err := builder.NewBuilder("gaming")
	if err != nil {
		panic(err)
	}
	director := builder.NewDirector(defaultBuilder)

	defaultComputer := director.BuildComputer()
	defaultComputer.DisplayInfo()

	director.SetBuilder(gamingBuilder)

	gamingComputer := director.BuildComputer()
	gamingComputer.DisplayInfo()
}
