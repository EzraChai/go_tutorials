package main

import "fmt"

type gasEngine struct {
	gallons int
	mpg     int
}

type electricEngine struct {
	kwh   float64
	mpkwh float64
}

type car[T gasEngine | electricEngine] struct {
	carModel string
	carMake  string
	engine   T
}

func main() {
	newCar := car[electricEngine]{
		carModel: "2025 Rangers",
		carMake:  "Ford",
		engine: electricEngine{
			kwh:   100,
			mpkwh: 100,
		},
	}

	fmt.Println(newCar)
}
