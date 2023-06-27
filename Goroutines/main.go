package main

import (
	"fmt"
	"time"
)

func main() {

	//In Golang we can launch a gorouine by 'go' keyword.
	//This function is called by launching of goroutine and result of this will be shown when goroutine will complete it's execution.
	go greeter("Harshvardhan")

	//This function is normally called
	greeter("Yash")
}

// this function will basically greet a person.
func greeter(s string) {

	for i := 0; i < 6; i++ {

		//This time will basically tell main() method to wait for this routine.
		//If we will not use this method than main mehtod will not wait for goroutine's task completion and we cant see "Hello Harshvardhan" printed.
		//There are also some pakages like 'sync' for synchronization which can resolve this issue. We will see it in Tutorial of Mutaxts & Wiatgroups.
		time.Sleep(30 * time.Millisecond)

		fmt.Printf("Hello %s\n", s)
	}

}
