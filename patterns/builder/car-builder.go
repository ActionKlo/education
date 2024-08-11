package main

import "fmt"

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type Car struct {
	color    Color
	wheels   Wheels
	topSpeed Speed
}

type carBuilder struct {
	car *Car
}

func NewBuilder() Builder {
	return &carBuilder{&Car{}}
}

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Interface
}

type Interface interface {
	Drive() error
	Stop() error
}

func (c carBuilder) Color(color Color) Builder {
	c.car.color = color
	return c
}

func (c carBuilder) Wheels(wheels Wheels) Builder {
	c.car.wheels = wheels
	return c
}

func (c carBuilder) TopSpeed(speed Speed) Builder {
	c.car.topSpeed = speed
	return c
}

func (c carBuilder) Build() Interface {
	return c.car
}

func (c Car) Drive() error {
	fmt.Println("drive", c)
	return nil
}

func (c Car) Stop() error {
	fmt.Println("stop", c)
	return nil
}

func main() {
	assembly := NewBuilder()

	familyCar := assembly.Color(RedColor).Wheels(SportsWheels).TopSpeed(50 * MPH).Build()
	_ = familyCar.Drive()
	_ = familyCar.Stop()

	assembly = NewBuilder()

	sportsCar := assembly.Wheels(SteelWheels).TopSpeed(150 * MPH).Build()
	_ = sportsCar.Drive()
	_ = sportsCar.Stop()
}
