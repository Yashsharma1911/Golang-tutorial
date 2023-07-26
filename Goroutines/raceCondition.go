package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race conadition")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One R")

		score = append(score, 1)

		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two R")
		score = append(score, 2)
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three R")
		score = append(score, 3)
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
