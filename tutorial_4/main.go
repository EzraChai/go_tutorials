package main

import "fmt"

type PetrolEngine struct {
	km_per_litre float64
	litre        float64
}

type ElectricEngine struct {
	km_per_kwh float64
	kwh        float64
}

func (e ElectricEngine) km_left() float64 {
	return e.km_per_kwh * e.kwh
}

func (e PetrolEngine) km_left() float64 {
	return e.km_per_litre * e.litre
}

type Engine interface {
	km_left() float64
}

func canMakeIt(engine Engine, km float64) bool {
	return engine.km_left() > km
}

//------------------------------------------------------------------------------

type Student struct {
	Name       string
	Age        int
	University University
}

type People interface {
}

type University struct {
	Name string
}

func main() {
	// var uni1 University = University{Name: "Universiti Malaya"}
	// var student1 Student = Student{Name: "Ezra Chai", Age: 19, University: uni1}
	// fmt.Println(student1)
	// printStudentName(student1)
	// fmt.Println(student1.Name + "'s university is " + student1.University.Name)

	// fmt.Println(uni1)

	var en1 PetrolEngine = PetrolEngine{km_per_litre: 100, litre: 10}
	var en2 ElectricEngine = ElectricEngine{km_per_kwh: 100, kwh: 10}
	fmt.Println(canMakeIt(en1, 100))
	fmt.Println(canMakeIt(en2, 100))
}
