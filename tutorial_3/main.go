package main

import "fmt"

func main() {
	var name []rune = []rune("张三")
	fmt.Println(name)
	for i, v := range name {
		fmt.Println(i, v)
	}
}
