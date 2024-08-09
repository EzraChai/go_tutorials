package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	num := []int{1, 2, 4, 5, 6}
	result := squareEveryNumber(num)
	fmt.Printf("%p\n", &result)
	fmt.Printf("%p\n", &num)
	fmt.Printf("%v\n", result)
	fmt.Printf("%v\n", num)
	fmt.Println(time.Since(currentTime))
}

func squareEveryNumber(numbers []int) []int {
	for i := range numbers {
		(numbers)[i] = (numbers)[i] * (numbers)[i]
	}
	return numbers
}
