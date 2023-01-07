<span align="center">
 <h1>File management in Golang</h1>
</span>

## What is File management in Golang?
**File management** in Golang involves using functions and methods provided by various packages to perform operations on files and directories, such as reading and writing file contents, creating and deleting files and directories, and renaming files and directories. Some important packages for file management in Golang include `os`, `ioutil`, `bufio`, `encoding/csv`, and `encoding/json`. These packages provide functions and methods for working with files and directories, reading and writing file contents, and parsing and generating common file formats such as CSV and JSON. *Let's explore some common packages for Golang 👇*

## Why we use OS package ?

This package provides functions for interacting with the operating system. It includes functions for opening, closing, reading, and writing files, as well as functions for creating, deleting, and renaming files and directories.

## Function use cases of OS package

### 1. os.Open()
This function opens a file and returns a *File value that can be used to read or write the file. The function takes the file name and the file mode
```
f, err := os.Open("filename.txt")
if err != nil {
    log.Fatal(err)
}
defer f.Close()
```

### 2. os.Create()
This function creates a new file or overwrites an existing file. It takes the file name as an argument and returns a *File value that can be used to write to the file.
```
f, err := os.Create("filename.txt")
if err != nil {
    log.Fatal(err)
}
defer f.Close()
```

### 3. os.Mkdir()
This function creates a new directory. It takes the directory name as an argument and returns an error if it fails to create the directory.
```
err := os.Mkdir("mydir", 0755)
if err != nil {
    log.Fatal(err)
}
```

### 4. os.Remove()
This function deletes a file or directory. It takes the name of the file or directory as an argument and returns an error if it fails to delete the file or directory.
```
err := os.Remove("filename.txt")
if err != nil {
    log.Fatal(err)
}
```

### 5. os.Rename()
This function renames a file or directory. It takes the current name and the new name as arguments and returns an error if it fails to rename the file or directory.
```
err := os.Rename("oldname.txt", "newname.txt")
if err != nil {
    log.Fatal(err)
}
```
## Explore other packages

- [Ioutil package](https://github.com/Yashsharma1911/Golang-tutorial/blob/main/File/common-packages/ioutil_package.md)
