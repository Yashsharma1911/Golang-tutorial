## Why we use ioutil package ?

The ioutil package in Go provides a number of functions for working with input/output (I/O) operations. It includes functions for reading and writing files, as well as functions for working with in-memory buffers that can be used to store data.

Here are a few examples of functions you can find in the ioutil package:

1. `ReadFile`: reads the entire contents of a file and returns it as a slice of bytes.
2. `WriteFile`: writes data to a file. If the file does not exist, it will be created. If it does exist, it will be overwritten.
3. `ReadDir`: reads the contents of a directory and returns a slice of os.FileInfo values, one for each file or subdirectory in the directory.
   TempFile: creates a new temporary file and returns a \*os.File pointing to it.
