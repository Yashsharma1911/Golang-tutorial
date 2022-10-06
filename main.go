package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")

	// comma ok idiom
	text, _ := reader.ReadString('\n')
	fmt.Println("Thanks", text)
}
