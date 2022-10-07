package main

// import these packages
import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("hello")

	// created a variable which will hold the data to put in file
	content := "This is how you can create file in GO"

	// Create a file
	// creation of file is OS operation that's why it uses os package
	file, err := os.Create("./test.txt")
	// made function to check if error is nil or not
	checkNilErr(err)

	// Write to file and you can check length of bytes written
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("file length: ", length)

	// Close file
	// prefer to use defer keyword on closing file
	defer file.Close()

	// Created function to read file data
	readFile("./test.txt")
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(path string) {
	// Read file
	// ioutil package is used to read file

	databyte, err := ioutil.ReadFile(path)
	checkNilErr(err)              // check error
	fmt.Println(string(databyte)) // print data
}
