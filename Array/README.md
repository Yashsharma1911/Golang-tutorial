<span align="center">
 <h1>Array in Golang</h1>
</span>
<br>
In Go, an array is a fixed-size sequence of elements of the same type. The size of an array is part of its type, so the types [10]int and [20]int are distinct.

Here is the syntax for declaring an array in Go:

```
var arr [size]type
```

<br>

For example, to declare an array of 10 integers, you would write:

```
var arr [10]int
```

<br>

You can also use an initializer to specify the values of the array elements when you declare the array. For example:

```
arr := [5]int{1, 2, 3, 4, 5}
```

<br>

This declares an array arr of 5 integers and initializes the elements with the values 1, 2, 3, 4, and 5.

You can access the elements of an array using an index, which starts at 0. For example, the expression arr[i] accesses the i-th element of arr.

Here is a complete example that demonstrates how to declare, initialize, and access an array in Go:

```
package main

import "fmt"

func main() {
  // Declare and initialize an array of 5 integers
  arr := [5]int{1, 2, 3, 4, 5}

  // Print the first element of the array
  fmt.Println(arr[0])

  // Modify the third element of the array
  arr[2] = 100

  // Print the third element of the array
  fmt.Println(arr[2])
}
```

<br>
Output:

```
1
100
```
