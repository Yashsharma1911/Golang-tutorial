```markdown
<div align="center">
  <h1>Conversions in Go Lang</h1>
  <p>This readme contains examples and explanations of various types of conversions in the Golang. Conversions are an essential aspect of programming, enabling the transformation of data between different types to ensure compatibility and efficient processing.</p>
</div>

Explore the implementation examples and explanations provided in this repository to gain a better understanding of various conversion techniques in Go.

## Table of Contents

- [Introduction](#introduction)
- [Numeric Conversions](#numeric-conversions)
- [String Conversions](#string-conversions)
- [Type Assertion](#type-assertion)
- [Interface Conversions](#interface-conversions)
- [Custom Type Conversions](#custom-type-conversions)
- [Pointer Conversions](#pointer-conversions)
- [Array and Slice Conversions](#array-and-slice-conversions)
- [Map Conversions](#map-conversions)
- [Struct Conversions](#struct-conversions)
- [Time and Duration Conversions](#time-and-duration-conversions)

## Introduction

Conversions involve changing the type of a value to another type, either implicitly or explicitly. This repository provides examples and explanations for various scenarios where conversions are commonly used in Go.

## Numeric Conversions

Numeric conversions involve converting between different numeric types, such as `int`, `float64`, `uint32`, etc. Learn about type casting, precision considerations, and potential data loss.

### Implementation Example:

```go
package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b float64 = float64(a)

	fmt.Printf("a: %d\n", a)
	fmt.Printf("b: %f\n", b)
}
```

## String Conversions

Learn how to convert between strings and other data types, including numbers and non-ASCII characters. Understand encoding and decoding processes for strings.

### Implementation Example:

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Converted number: %d\n", num)
}
```

## Type Assertion

Type assertions allow you to determine the underlying type of an interface and extract its value in the correct type. Explore how to use type assertions effectively.

### Implementation Example:

```go
package main

import (
	"fmt"
)

func main() {
	var val interface{}
	val = 42

	if num, ok := val.(int); ok {
		fmt.Printf("Value is an integer: %d\n", num)
	} else {
		fmt.Println("Value is not an integer")
	}
}
```

## Interface Conversions

Understand how to work with interfaces and convert between different interface types. Learn about empty interfaces and type switches.

### Implementation Example:

```go
package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	var s Shape
	s = Circle{Radius: 5.0}

	circle, ok := s.(Circle)
	if ok {
		fmt.Printf("Circle Area: %f\n", circle.Area())
	} else {
		fmt.Println("Conversion failed")
	}
}
```

## Custom Type Conversions

Explore how to define your own custom types and perform conversions between these types and built-in types. Learn about the importance of method receivers in type conversions.

### Implementation Example:

```go
package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func main() {
	celsius := Celsius(30)
	fahrenheit := celsius.ToFahrenheit()

	fmt.Printf("%.2f°C is %.2f°F\n", celsius, fahrenheit)
}
```

## Pointer Conversions

Pointer conversions involve converting between different pointer types. Learn about pointer dereferencing, type compatibility, and common pitfalls.

### Implementation Example:

```go
package main

import (
	"fmt"
)

type Animal struct {
	Name string
}

func main() {
	cat := Animal{Name: "Cat"}
	var animalPtr *Animal
	animalPtr = &cat

	fmt.Printf("Name: %s\n", (*animalPtr).Name)
}
```

## Array and Slice Conversions

Arrays and slices can be converted between each other, and this section explains how to perform such conversions efficiently.

### Implementation Example:

```go
package main

import (
	"fmt"
)

func main() {
	arr := [3]int{1, 2, 3}
	slice := arr[:]

	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Slice: %v\n", slice)
}
```

## Map Conversions

Learn about converting between different map types and techniques for copying or merging map data.

### Implementation Example:

```go
package main

import (
	"fmt"
)

func main() {
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"b": 3, "c": 4}

	for key, value := range map1 {
		map2[key] = value
	}

	fmt.Printf("Merged Map: %v\n", map2)
}
```

## Struct Conversions

Structs can be converted between different types that share the same underlying field structure. Explore techniques for converting between related struct types.

### Implementation Example:

```go
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Name   string
	Salary float64
}

func main() {
	person := Person{Name: "Alice", Age: 30}
	employee := Employee(person)

	fmt.Printf("Employee: %+v\n", employee)
}
```

## Time and Duration Conversions

Converting time values and durations between different formats and units is essential when working with time-sensitive applications. Learn how to handle time conversions effectively.

### Implementation Example:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("Current time: %v\n", now)

	utcTime := now.UTC()
	fmt.Printf("UTC time: %v\n", utcTime)

	layout := "2006-01-02 15:04:05"
	timeStr := "2023-08-18 12:30:00"
	parsedTime, _ := time.Parse(layout, timeStr)
	fmt.Printf("Parsed time: %v\n", parsedTime)
}
```

## Explore other packages

- [Ioutil package](https://github.com/Yashsharma1911/Golang-tutorial/blob/main/File/common-packages/ioutil_package.md)

---
```
