<span align="center">																													
<h1>if-else statement in Golang</h1>
</span>

## Why do we use if statement in Go?
The " if " statement is used to check if a condition is fulfilled and executes the block of code only when the condition is met.

syntax for writing an if statement : 
```go
if condition{
         //block of code to be executed if the condition is met 
}
```

We can also initialize a variable and check the condition within the if statement :
```go
if <init>; condition{
	//block of code to be executed when the condition is true.
}
```

## Why do we use else statement in Go?
The " else " statement is used when we have to execute a block of code if the condition in the " if " statement is not met.

syntax for writing an else statement:
```go
if condition{
	//block of code to be executed if the condition is met
}else{
	//block of code to be executed if the condition is not met
}
```
**note** : The else statement has to be in the same line where the if statement was closed with " } "

## When do we use else if statement in Go?
We can cascade an if statement in an else statement in goLang , to check for more conditions . It is generally called as else if statements .

syntax for writing an else if statement:
```go
if condition1{
	//block of code to be executed when condition 1 is met
}else if condition2{
	//block of code to be executed when condition 2 is met
}else{
	// block of code to be executed when no condition is met
}
```

## Using if-else with boolean experssions :
Along with conditions we can also use boolean expressions in Go :
```go
if boolean1{
	// code will be executed if the boolean1 is true 
}else if !boolean2{
	//code will be executed if the boolean2 is false
}else{
	//code to be executed when no above condition is met 
}
```
We know that any boolean experssion has only 2 values i.e. true or false . When we use boolean experssions with if-else statement, The bool variable with value true means the condition is met and the corresponding block of code inside the statement will be executed.


## Some important points to keep in mind while using if-else statement:

* an " if " statement can have zero or one " else " statements but it has to be after all the " else if " statements
* once an " else if " condition is met , none of the remaining " else if " or " else "statements are going to be checked .
