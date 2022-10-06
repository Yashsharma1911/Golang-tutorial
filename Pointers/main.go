package main

import "fmt"

func main() {
	var myInt2 int
	myInt2 = 42

	myInt := &myInt2
	myInt3 := &myInt

	fmt.Println("Value of myInt is:", **myInt3)
}
