package main

import (
	"fmt"
)

func main() {
	var fruit [3]string
	fruit[0] = "apple"
	fruit[2] = "banana"
	fmt.Println(len(fruit))
}
