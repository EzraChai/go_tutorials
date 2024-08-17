package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return math.Pow(s.length, 2)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func findArea(s shape) {
	fmt.Printf("Area: %.4f\n", s.area())
}

func main() {
	fmt.Println("Hello World")
	c1 := circle{
		radius: 2,
	}
	s1 := square{
		length: 2,
	}
	findArea(c1)
	findArea(s1)
}
