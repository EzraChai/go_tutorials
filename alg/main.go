package main

import "fmt"

func findSameDigits(arr []int) int {
	var map1 map[int]bool = make(map[int]bool)
	for i := range arr {
		if !map1[arr[i]] {
			map1[arr[i]] = true
		} else {
			fmt.Println(map1)
			return arr[i]
		}
	}
	return 0
}

func main() {
	arr := []int{1, 6, 3, 4, 5, 6, 8, 9, 10}
	fmt.Println(findSameDigits(arr))
}
