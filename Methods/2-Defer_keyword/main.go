package main

import "fmt"

// Defer keyword
//
// The defer keyword is used to delay the execution of a function until the surrounding function returns.

// The defer works on first in last out basis.

// The deferred function executes even if a run-time panic occurs.

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

// Output:
// hello
// world
