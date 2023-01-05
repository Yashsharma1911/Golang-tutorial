<span align="center">
 <h1>Hello World 👋</h1>
</span>
<br>

## What is Fmt package ?

The fmt package is a standard library package in Go that provides functions for formatted I/O operations. It includes functions for printing to the console, reading from and writing to files, and formatting strings.

Here are some of the common functions that you can use from the fmt package:

1. `Print()`, `Println()`, `Printf()`: These functions are used to print output to the console.

2. `Scan()`, `Scanf()`, `Scanln()`: These functions are used to read input from the console.

3. `Fprint()`, `Fprintln()`, `Fprintf()`: These functions are similar to the Print family of functions, but they allow you to specify an output stream (such as a file) to which the output should be written.

4. `Sprint()`, `Sprintln()`, `Sprintf()`: These functions are similar to the Print family of functions, but they return the formatted string instead of printing it to the console.

5. `Errorf()`: This function is used to print an error message to the console.

## Some Use cases

1. `fmt.Print()`: This function prints a string to the console, followed by a newline.

```
// Print a string to the console
fmt.Print("Hello, world!\n")
```

<br>

2. `fmt.Println()`: This function is similar to fmt.Print(), but it adds a newline at the end of the output.

```
// Print a string with a newline at the end
fmt.Println("Hello, world!")
```

<br>

3. `fmt.Printf()`: This function allows you to use a format string to specify how the output should be formatted. You can use this function to print values of different types, such as integers, floating-point numbers, and strings.

```
// Print a floating-point number
fmt.Printf("%.2f", 3.14159)
```

<br>

4. `fmt.Sprintf()`: This function is similar to fmt.Printf(), but it returns the formatted string instead of printing it to the console.

```
// Use Sprintf to create a formatted string
s := fmt.Sprintf("%s is %d years old", "Alice", 30)
fmt.Println(s)
```

<br>
