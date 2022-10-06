package main

// functions
// function is a block of code that is used to perform a specific task

// syntax
/*
	func functionName(parameterName type) returnType {
		// code
	}
*/

// function with no parameter and no return type
/*
	func hello() {
		fmt.Println("Hello World")
	}
*/

// function with parameter and no return type
/*
	func hello(name string) {
		fmt.Println("Hello", name)
	}
*/

// function with no parameter and return type
/*
	func hello() string {
		return "Hello World"
	}
*/

// function with parameter and return type
/*
	func hello(name string) string {
		return "Hello " + name
	}
*/

// function with multiple parameter and return type
/*
	func hello(name string, age int) string {
		return "Hello " + name + " you are " + strconv.Itoa(age) + " years old"
	}
*/

// function with multiple parameter and multiple return type
/*
	func hello(name string, age int) (string, string) {
		return "Hello " + name + " you are " + strconv.Itoa(age) + " years old", "Hello " + name
	}
*/

// function with parameter as an slice
/*
	func hello(values ...int) int {
		total := 0
		for _, val := range values {
			total += val
		}
		return total
	}
*/

func main() {
	// hello()
	// hello("Yash")
	// fmt.Println(hello())
	// fmt.Println(hello("Yash"))
	// fmt.Println(hello("Yash", 21))
	// fmt.Println(hello(1, 2, 3, 4, 5, 6, 7, 8))
}
