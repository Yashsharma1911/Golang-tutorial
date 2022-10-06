package main

import "fmt"

func main() {
	// for loop
	// for init; condition; post {}
	// init is executed only once at the beginning of the loop
	// condition is evaluated at the beginning of every iteration
	// post is executed at the end of every iteration

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// for loop with condition
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for loop without condition
	// i := 0
	// for {
	// 	if i >= 10 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }

	// for loop with range
	// for index, value := range "yash" {
	// 	fmt.Println(index, string(value))
	// }

	// for loop with range
	// for index, value := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
	// 	fmt.Println(index, value)
	// }

	// for loop with range
	// for index, value := range map[string]string{"a": "apple", "b": "ball"} {
	// 	fmt.Println(index, value)
	// }

	// for loop with range
	// for index, value := range map[int]string{1: "apple", 2: "ball"} {
	// 	fmt.Println(index, value)
	// }

	// for loop with range
	// for index, value := range map[string]int{"a": 1, "b": 2} {
	// 	fmt.Println(index, value)
	// }

	// for loop with range
	// for index, value := range map[int]int{1: 1, 2: 2} {
	// 	fmt.Println(index, value)
	// }

	// for loop with range
	// for index, value := range map[int]string{1: "apple", 2: "ball"} {
	// 	fmt.Println(index, value)
	// }

	// for loop with range
	// for index, value := range map[int]string{1: "apple",
	fruits := []string{"apple", "banana", "orange", "grape", "mango", "pineapple", "watermelon", "cherry"}

	// for loop
	for i := 0; i < 5; i++ {
		// fmt.Println(i)
	}

	// for loop with range
	for index, value := range fruits {
		fmt.Println("index:", index, "value:", value)
	}

	// while loop
	j := 0
	for j < 5 {
		// fmt.Println(j)
		j++
	}

	// infinite loop
	k := 0
	for {
		// fmt.Println(k)
		k++
		if k == 5 {
			break
		}
	}
}
