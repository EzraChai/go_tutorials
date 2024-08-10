package main

import "fmt"

func main() {
	var channel = make(chan int)
	go process(channel)
	for v := range channel {
		fmt.Println(v)
	}
}

func process(channel chan int) {
	defer close(channel)

	for i := 0; i < 10; i++ {
		channel <- i
	}
}
