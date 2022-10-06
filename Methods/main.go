package main

import "fmt"

// Methods

// Methods are functions with a special receiver argument
// the method contains a receiver argument in it.
// With the help of the receiver argument, the method can access the properties of the receiver.
// Here, the receiver can be of struct type or non-struct type.

// syntax
/*
	func (receiverName receiverType) methodName(parameterName parameterType) returnType {
		// code
	}
*/

func (u User) getName() {
	fmt.Println(u.Name)
}

func main() {
	client := User{"yash", "yashsharma2572@gmail.com", 20}
	client2 := User{"sudhanshu", "sudhanshu2@gmail.com", 21}

	client.getName()
	client2.getName()
}

// struct
type User struct {
	Name  string
	Email string
	Age   int
}
