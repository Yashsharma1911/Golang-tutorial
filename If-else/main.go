package main

import "fmt"

func main() {
	// if else syntax

	// if condition {
	// 	// do something
	// } else if condition {
	// 	// do something
	// } else {
	// 	// do something
	// }

	// through variable
	var value int = 34

	// you need to follow the below syntax
	if value > 10 {
		fmt.Println("value is greater then 10")
	} else if value < 10 {
		fmt.Println("value is less then 10")
	} else {
		fmt.Println("value is less then 10")
	}

	// if condition
	// {
	// this will show error!
	// }

	// we can initialize the variable inside the if condition
	if value := 9; value > 10 {
		fmt.Println("value is greater then 10")
	} else {
		fmt.Println("value is less then 10")
	}
}
