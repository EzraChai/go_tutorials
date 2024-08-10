package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice2 := []float64{1.1, 2.2, 3.6, 4.5, 5.2}

	num := sumSlice[int](slice)
	fmt.Printf("Sum = %v\n", num)

	num2 := sumSlice[float64](slice2)
	fmt.Printf("Sum = %v\n", num2)
}

// Generics
func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}
