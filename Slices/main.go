package main

import (
	"fmt"
	"sort"
)

func main() {
	// slice is a data structure that is used to store a list of data
	// slice is a dynamic array
	// you don't need to specify the size of the slice in advance like array in golang

	// syntax
	mySlice := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	// to add value in slice
	mySlice = append(mySlice, "i")

	// print value how much you need
	// 1 is starting value to cut from and 6 is ending value to cut
	fmt.Println(mySlice[1:6])

	// to sort the slice
	mySliceInteger := []int{12, 45, 456, 782, 8, 5415}
	sort.Ints(mySliceInteger)
	fmt.Println(mySliceInteger)

	// to sort the slice
	nameList := []string{"John", "Paul", "George", "Ringo", "yash", "sharma", "sudhanshu"}
	sort.Strings(nameList)

	// sclice your slice with append function to get what you want
	nameList = append(nameList[:2], nameList[4:]...) // output: [George John sharma sudhanshu yash]
	fmt.Println(nameList)
}
