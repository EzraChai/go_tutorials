package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup = sync.WaitGroup{}
var m = sync.RWMutex{}
var DBData []int = []int{1, 2, 3, 4, 5}
var results []int = []int{}

func main() {
	timeNow := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go DBCall()
	}
	wg.Wait()
	fmt.Println("Time taken to execute all the DB calls: ", time.Since(timeNow))

}

func DBCall() {
	i := rand.Intn(6)
	time.Sleep(time.Duration(500) * time.Millisecond)
	save(i)
	log()
	fmt.Println("The result from DB is ", i)
	fmt.Println("Total number of results: ", results)
	wg.Done()
}

func save(i int) {
	m.Lock()
	results = append(results, i)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Println("The current results are ", results)
	m.RUnlock()
}
