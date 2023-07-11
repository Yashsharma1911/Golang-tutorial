package main

import (
	"fmt"
	"sync"
)

var mut sync.Mutex    //should be pointer
var wg sync.WaitGroup //should be pointer
var list = []string{"Hello World"}

func main() {
	go appendString("How are you")
	go appendString("Hi There")
	go appendString("Hola")

	/*
		we are launching multiple goroutine at same time so there is possiblity that two or more gorotine will get enter in
		critical section and we can get overriden data.

		This problem can be solved by using mutex.
	*/

	wg.Add(3)

	wg.Wait()
}

func appendString(s string) {
	//here list is a critical section. multiple go routines are going to use it at same time to so we have to proctct it.
	//for doing this we use mutext. We put a lock when a goroutine comes in this section and will unlock when task is completed.
	mut.Lock()
	list = append(list, s)
	mut.Unlock()

	fmt.Println(list)
	defer wg.Done()
}
