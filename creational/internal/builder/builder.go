package builder

import (
	"errors"
	"fmt"
)

// Product
type Computer struct {
	cpuName string
	ramInfo string
	storage int
}

func (c *Computer) setCPU(newCpuName string) {
	c.cpuName = newCpuName
}

func (c *Computer) setRAM(newRamInfo string) {
	c.ramInfo = newRamInfo
}

func (c *Computer) setStorage(newStorage int) {
	c.storage = newStorage
}

func (c *Computer) DisplayInfo() {
	fmt.Printf("<Computer congifuration>\nCPU: %s\nRAM: %s\nDisk: %d Gb\n",
		c.cpuName, c.ramInfo, c.storage)
}

// Builders Interface
type BuilderInterface interface {
	buildCPU()
	buildRAM()
	buildStorage()
	getResult() *Computer
}

// Builder for default computers
type DefaultBuilder struct {
	comp *Computer
}

func (b *DefaultBuilder) buildCPU() {
	b.comp.setCPU("Intel Core i3-12100")
}

func (b *DefaultBuilder) buildRAM() {
	b.comp.setRAM("8GB DDR4")
}

func (b *DefaultBuilder) buildStorage() {
	b.comp.setStorage(500)
}

func (b *DefaultBuilder) getResult() *Computer {
	return b.comp
}

// Builder for gaming computers
type GamingBuilder struct {
	comp *Computer
}

func (b *GamingBuilder) buildCPU() {
	b.comp.setCPU("Intel Core i9-13900KS")
}

func (b *GamingBuilder) buildRAM() {
	b.comp.setRAM("32GB DDR5")
}

func (b *GamingBuilder) buildStorage() {
	b.comp.setStorage(2000)
}

func (b *GamingBuilder) getResult() *Computer {
	return b.comp
}

// Method for creating builder
func NewBuilder(builderType string) (BuilderInterface, error) {
	if builderType == "default" {
		return &DefaultBuilder{comp: &Computer{}}, nil
	} else if builderType == "gaming" {
		return &GamingBuilder{comp: &Computer{}}, nil
	} else {
		return nil, errors.New("invalid type of builder")
	}
}

// Director for configurating process of building
type Director struct {
	builder BuilderInterface
}

func (d *Director) SetBuilder(b BuilderInterface) {
	d.builder = b
}

func (d *Director) BuildComputer() *Computer {
	d.builder.buildCPU()
	d.builder.buildRAM()
	d.builder.buildStorage()
	return d.builder.getResult()
}

func NewDirector(builder BuilderInterface) Director {
	return Director{builder: builder}
}
