package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5

func main() {
	channel := make(chan string)
	website := []string{"costco", "amazon", "walmart", "target"}
	for _, v := range website {
		go findPrice(v, channel)
	}
	sendMessage(channel)
}

func sendMessage(channel chan string) {
	fmt.Println("Found a deal on chicken at: ", <-channel)
}

func findPrice(website string, channel chan string) {

	for {
		time.Sleep(time.Second)
		var chickenPrice float32 = rand.Float32() * 20
		if chickenPrice < MAX_CHICKEN_PRICE {
			channel <- website
			break
		}
	}
}
