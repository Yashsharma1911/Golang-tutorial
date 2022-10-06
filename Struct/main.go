package main

import "fmt"

func main() {
	// In golang there is no inheritance and no class and no super and parent
	// but we can use struct to create our own type

	// type myStruct struct {
	// 	name string
	// 	age  int
	// }

	// var myStructVar myStruct
	// myStructVar.name = "yash"
	// myStructVar.age = 23
	// fmt.Println(myStructVar)

	client := User{"yash", "yashsharma@gmail.com", 23}
	fmt.Println(client)
	fmt.Printf("%+v", client) // for more detail
	fmt.Println(client.Name)  // fir accessing the perticular value of struct
}

// you can create struct outside
type User struct {
	Name  string
	Email string
	Age   int
}
