<span align="center">
 <h1>Maps in Golang</h1>
</span>

#### Contributor : [Sagar Singh](https://github.com/SagarSingh2003)

## What are Maps?
Map is a very important data structure in Golang used to store ,add and delete data in your program. Go provides a built-in type map that implements a hash table. 
### Uses :
- Provides fast lookup 
- Store key value pairs and does not allow duplicate key values

## How to declare and initialize maps in Golang?

- ### Declaring a map type:
  ```go
  var map1 map[<key type>]<valuetype>
  ```
  here map1 is the variable name that represents the memory adress where the map is stored.
  
  ```go
  eg:  var map1 map[string]int
  ```
  here the "key" type is string and the "value" type is integer.
  Map types are reference types like pointers and slice , so the above value of map1 is nil , It does not point to an initialized map.
  And writing attempt to a nil map will cause a runtime Panic, so don't do that.
  

  Here is an example :

  ![02img](https://github.com/SagarSingh2003/Golang-tutorial/assets/129133613/47a521e1-f2a7-4e8f-b733-c1c78961f279)
  ![03img](https://github.com/SagarSingh2003/Golang-tutorial/assets/129133613/615328aa-b187-4af7-920b-a479c6da80ca)

- ### Initializing a map:
  To initialize a map we use make function
  ```go
  map1 = make(map[string]int, 100)
  ```
  Here , 100 is capacity of the map that the user has to set in order to initialize the map.
  But the initial capacity does not bound it's size , i.e. when more than a hundered elements will be added to the map, it will expand accordingly. 
  This means that the map grows to accomodate the size of the items stored in it, with an exception of nil map because items can't be added in it.

  To initialize a map with some data 
  ```go
  map1 := map[string]int{
  "yash": 1,
  }
  ```

## Adding , deleting and formatting data in a map:
  To add a value in an initialized map
  ```go
  map1["sagar"] = 2
  ```

  To retrieve data from a map we use keys to access the values , here the key "sagar" has the value 2 which is now stored in the variable "i"
  ```go
  i := map1["sagar"]
  ```

  To delete data from a map we use the inbuilt delete function , here we want to remove the key value pair " sagar: 2 " from map1
  ```go
  delete(map1, "sagar")
  ```
  
  To interate over the contents of a map, we use range keyword:
  ```go
  for key, value := range map1{
     fmt.Println("key:", key,"value:",value)
  }
  ```

  if we don't want the key or the value we swap it with underscore or else we get error : declared but not used
  ```go
  for _, value := range map1{
    fmt.Println("value:",value)
  }
  ```

## Checking if a key exists or not using "comma ok" syntax
  Here "i" is the variable that stores the value of the key "yash" , "ok " is of type bool that has value true if the key exists  
  ```go
  i, ok := map1["yash"]
  ```
  If we don't want to store the value we use an underscore(_)
  ```go
  _, ok := map1["yash"]
  ```
  
## Some important points :
  - if a requested key is not present then we get the "value" type's zero value 
    ```go
    j := map1["kunal"]
    //j == 0
    ```
    kunal key does not exist so, we get zero value of type integer that is "0"

  - To find out the numer of elements in the map we use len function 
    ```go
    k := len(map1)
    fmt.Println(k)
    ```
    output 
    ```go
    1
    ```
    there is only one key value pair in map1 "yash: 1" , so we get k = 1

## For learning more about maps :
   - [go documentation](https://go.dev/blog/maps)
