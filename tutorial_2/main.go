package main

import (
	"fmt"
	"time"
)

func main() {
	var intArr [3]int
	intArr[0] = 1
	intArr[1] = 2
	intArr[2] = 3
	fmt.Println(intArr[1:3])
	for i := 0; i < 3; i++ {
		println(&intArr[i])
	}

	intArr2 := [...]int{1, 2, 3}
	fmt.Println(intArr2)

	var intSlice = []int{1, 2, 3}
	fmt.Println(len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 4)
	intSlice = append(intSlice, 5)
	fmt.Println(intSlice)
	fmt.Println(len(intSlice), cap(intSlice))

	var intSlice2 = make([]int32, 3, 8)
	fmt.Println(intSlice2)

	myMap := make(map[string]uint8)
	fmt.Println(myMap)
	myMap["count"] = 1
	fmt.Println(myMap["count"])
	myMap["age"] = 10
	delete(myMap, "age")
	var age, ok = myMap["age"]
	if ok {
		fmt.Println(age)
	} else {
		fmt.Println("not found")
	}

	for index, val := range myMap {
		fmt.Printf("Index:%s, Value:%d\n", index, val)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	i := 0
	for i < 10 {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()

	time := timeLoop(func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d ", i)
		}
		fmt.Println()
	})
	fmt.Println(time)

	time2 := timeLoop(func() {
		i := 0
		for i < 10 {
			fmt.Printf("%d ", i)
			i++
		}
		fmt.Println()
	})
	fmt.Println(time2)

}

func timeLoop(fun func()) time.Duration {
	start := time.Now()
	fun()
	return time.Since(start)
}
