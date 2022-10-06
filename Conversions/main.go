package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter rating between 1 to 5: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating us", input)

	newRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("add one you your string", newRating+1)
	}
}
