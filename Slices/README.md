
# Understanding Slices in Go

Slices are like dynamic lists that allow you to manage collections of items in Go. They're super useful because they're flexible, efficient, and can change in size. Think of a slice like a window into an array, where you can look at just a part of it, and even modify it!
## What's a Slice?

A slice is a way to handle a bunch of items in Go. It's like a bag you can put stuff in and take stuff out of. Imagine you have a bunch of numbers, and you want to do cool things with them. Slices let you do that without worrying too much about how big the bag is.




## Creating Slices

Making a slice is easy. You can start with an empty one or turn a regular list into a slice. Here's how:

```
// Creating an empty slice
mySlice := []int{}

// Turning a list into a slice
myList := [5]int{1, 2, 3, 4, 5}
mySlice := myList[1:4]  // Gets elements 2, 3, 4

```

## Slicing and Dicing
Slices let you pick out parts of your list easily. It's like you're slicing a cake into pieces. You just say where you want to start and where you want to stop, and you get a new slice with those elements.
```
myFruits := []string{"apple", "banana", "cherry", "date"}
mySlice := myFruits[1:3]  // Gets "banana" and "cherry"

```

## Adding and Growing
Slices are smart. You can add new things to them and they'll take care of growing themselves if needed. Just use the append function:
```
mySlice := []int{1, 2, 3}
mySlice = append(mySlice, 4, 5)  // Adds 4 and 5 to the slice


```
## Changing Things
Remember, slices are like windows into arrays. If you change something in the slice, you're also changing the original array. They're like magic mirrors! Look:

```
mySlice := []string{"hello", "world"}
mySlice[0] = "hi"
fmt.Println(mySlice)  // Prints: [hi world]

```

## resourses

Here are some related stuff

[YouTube video on slices By Hitesh Choudhary](https://youtu.be/k7hVj8QL9Co)

[Blog by golang official devlopers](https://go.dev/blog/slices-intro)
