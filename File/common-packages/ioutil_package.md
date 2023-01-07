## Why we use ioutil package ?

The ioutil package in Go provides a number of functions for working with input/output (I/O) operations. It includes functions for reading and writing files, as well as functions for working with in-memory buffers that can be used to store data.

Here are a few examples of functions you can find in the ioutil package:

1. `ReadFile`: reads the entire contents of a file and returns it as a slice of bytes.
2. `WriteFile`: writes data to a file. If the file does not exist, it will be created. If it does exist, it will be overwritten.
3. `ReadDir`: reads the contents of a directory and returns a slice of os.FileInfo values, one for each file or subdirectory in the directory.
   TempFile: creates a new temporary file and returns a \*os.File pointing to it.

## Function use cases

### 1. ioutil.ReadFile
reads the entire contents of a file and returns it as a slice of bytes.
```
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Read the contents of a file
	data, err := ioutil.ReadFile("/path/to/file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the contents of the file
	fmt.Println(string(data))
}

```

### 2. ioutil.WriteFile
IIn this example, we call the `ioutil.WriteFile` function and pass it the path to the file we want to write, a slice of bytes containing the data to write, and the file mode. The file mode controls the permissions of the file. In this case, the file mode is 0644, which means that the file will be readable and writable by the owner, and readable by everyone else.
```
package main

import (
	"io/ioutil"
)

func main() {
	// Write some data to a file
	err := ioutil.WriteFile("/path/to/file.txt", []byte("Hello, World!"), 0644)
	if err != nil {
		// Handle error
		return
	}
}
```

You can also use the ioutil.WriteFile function to append data to an existing file by passing the os.O_APPEND flag as the fourth argument. For example:
```
err := ioutil.WriteFile("/path/to/file.txt", []byte("Hello, World!"), 0644, os.O_APPEND)
```

### 3. ioutil.ReadDir
The ioutil.ReadDir function in Go reads the contents of a directory and returns a slice of os.FileInfo values,
```
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Read the contents of a directory
	files, err := ioutil.ReadDir("/path/to/directory")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Iterate over the slice of os.FileInfo values
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
```
In this example, we call `ioutil.ReadDir` and pass it the path to the directory we want to read. If the operation is successful, it will return a slice of `os.FileInfo` values and a **nil error**. If an error occurs, it will return a non-nil error value.

We can then use a for loop to iterate over the slice of os.FileInfo values and print the name of each file or subdirectory in the directory.

**Note** that the `os.FileInfo` type has many other fields that you can access, such as the size of the file, the last modification time, and the file mode (permissions). You can see a full list of the fields in the documentation for the os.FileInfo type: https://golang.org/pkg/os/#FileInfo
