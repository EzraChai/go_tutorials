package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int8 = 127
	intNum = intNum + 1
	fmt.Printf("intNum = %d\n", intNum)

	var floatNum float32 = 127.0
	fmt.Printf("floatNum = %f\n", floatNum)

	fmt.Printf("Total Num = %f\n", floatNum+float32(intNum))

	var int1 int8 = 3
	var int2 int8 = 2

	var ans int8 = int1 / int2

	fmt.Printf("ans = %d\n", ans)

	var myString string = "Hello, \nworld!"
	fmt.Printf("%s\n", myString)

	var myString2 string = `Hello, 
	world!`
	fmt.Printf("%s\n", myString2)

	fmt.Printf("%d\n", utf8.RuneCountInString(myString2))

	var myRune rune = 'a'
	myRune = '\u0041'
	fmt.Println(myRune)
	fmt.Printf("%c\n", myRune)

	myText := "sd"
	fmt.Printf("%s\n", myText)
	const myConst string = "Hello Guys"
	fmt.Println(myConst)

	printMe()
	printText("I love Go!")
	var num int = add(1, 1)
	fmt.Printf("num = %d\n", num)

	result, remainder, err := intDivision(10, 2)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else if remainder == 0 {
		fmt.Printf("result = %d\n", result)
	} else {
		fmt.Printf("result = %d\n", result)
		fmt.Printf("remainder = %d\n", remainder)
	}

	switch {
	case err != nil:
		fmt.Printf("Error: %s\n", err.Error())
	case remainder == 0:
		fmt.Printf("result = %d\n", result)
	default:
		fmt.Printf("result = %d\n", result)
		fmt.Printf("remainder = %d\n", remainder)
	}

}

func printMe() {
	fmt.Println("ME")
}

func printText(text string) {
	for _, char := range text {
		fmt.Printf("%c", char)
	}
}

func add(a, b int) int {
	return a + b
}

func intDivision(numerator, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Denominator cannot be zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, err

}
