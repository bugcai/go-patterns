package creational

import (
	"strconv"
)

type Color string

const (
	BlueColor Color = "blue"
	GreeColor       = "green"
	RedColor        = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type CarBuilder interface {
	Color(Color) CarBuilder
	Wheels(Wheels) CarBuilder
	TopSpeed(int) CarBuilder
	Build() Car
}

type Car interface {
	Drive() string
	Stop() string
}

type carBuilder struct {
	color    Color
	wheels   Wheels
	topSpeed int
}

func (cb *carBuilder) Color(color Color) CarBuilder {
	cb.color = color
	return cb
}

func (cb *carBuilder) Wheels(wheels Wheels) CarBuilder {
	cb.wheels = wheels
	return cb
}

func (cb *carBuilder) TopSpeed(speed int) CarBuilder {
	cb.topSpeed = speed
	return cb
}

func (cb *carBuilder) Build() Car {
	return &car{
		topSpeed: cb.topSpeed,
		color:    cb.color,
		wheels:   cb.wheels,
	}
}

func NewCarBuilder() CarBuilder {
	return &carBuilder{}
}

type car struct {
	color    Color
	topSpeed int
	wheels   Wheels
}

func (c *car) Drive() string {
	return "Driving at speed: " + strconv.Itoa(c.topSpeed)
}

func (c *car) Stop() string {
	return "Stopping a " + string(c.color) + " " + string(c.wheels) + " car"
}
