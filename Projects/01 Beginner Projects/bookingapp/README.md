<span align = "center">
  
  <h1>Ticket Booking Application in Go</h1>
    
</span>

## Note: 
This project is just shared to help you get better in Golang ! The idea of this project belongs to [TechWorldWithNana](https://youtu.be/yyUHQIec83I) Youtube Channel and we don't claim to be it's owner. 

## To learn how to build this application in a video format you can go the link below , have fun  :

- [TechWorldWithNana- ticket booking cli app](https://youtu.be/yyUHQIec83I)

## Aim :

In this exercise we are going to create a ticket booking application for the command line interface.

### Features:
- Greets User
- takes user input
- stores data in a data structure
- retrieves stored data
- checks for the validity of the user input
- books the no. of tickets  user wants for the Gopher conference
- shows the remaining tickets available

### Let's Go with Golang :

1. We start by creating a main.go file in a folder named bookingapp and opening the folder in your desired code editor

2. declare the package main in the main.go file

   ![5](https://github.com/SagarSingh2003/newME-git/assets/129133613/e99f6e73-7a21-4d0f-b773-e05b3fdc7725)


   when we run our program this is where the execution of the program begins 

3. open the main.go file in an integrated terminal and write the command :

   ```go
   go mod init bookingapp4
   ```

   this command initialized the go.mod file in your directory and there all the dependencies of the program will be stored.

  
   ![1002](https://github.com/SagarSingh2003/newME-git/assets/129133613/fdcbadcb-fc93-49a9-93d5-c5958c06cdf8)


4. declare func main() in the main.go file 

   ![2](https://github.com/SagarSingh2003/newME-git/assets/129133613/4727f9c8-cdb0-4a90-afa2-36eecf715d5a)


5. initialize variables 
   "confname" - stores the conference name

   "totTickets" - stores total number of tickets 

   "remTicket" - stores remaining number of tickets

   ![3](https://github.com/SagarSingh2003/newME-git/assets/129133613/5fe0dcaa-e696-4fb7-b9a9-f08a32e93be0)

6. now we go ahead and create a function GreetUsers() in our main.go file , to do that :

   ![4](https://github.com/SagarSingh2003/newME-git/assets/129133613/4ccfbd03-4f32-484c-bce1-d5f327b976b4)

   - to add functionality in the function we first import package fmt that allows us to format text or to print it on a file or console
   - we use Printf function of the fmt package to print the greeting message 

   ![image](https://github.com/SagarSingh2003/newME-git/assets/129133613/1c836cc2-be6a-4162-aad5-9ed58efec23c)


7. now when the function is declared we need to call it in our main function like this :
   ```go
   GreetUsers()
   ```

   ![1003](https://github.com/SagarSingh2003/newME-git/assets/129133613/aeaa1ef4-1c5a-413a-8ad2-d6ea30fb4ae6)

   we have to keep in mind that if we declare a function and don't call it inside the function main , it will not be exectued . 

9. Next we are gonna create a function that takes user input and we will ask for :

   - firstname
   - lastname
   - no. of tickets the user want to book
   - email address

   so , first we create the function and specify what type of variable we want in return from it

   ```go
   func GetUserInput() (string , string , int , string){
      
   }
   ```
10. we will use the Scan function from the fmt package to take user input
    - we first need to declare a variable
      ```go
      var firstname string
      ```
    - then we need to use the Scan function and specify the pointer or the adress we want to store the input in.
      ```go
      fmt.Println("Enter firstname :")
      fmt.Scan(&firstname) //here we use & to give the memory address of the variable firstname.
      ```
    - now we take user input for firstname , lastname , NoOfTicket , email and using the return keyword we specify what the function will return us .
      
      ![image](https://github.com/SagarSingh2003/newME-git/assets/129133613/83ec12a1-0f7b-49dd-911b-26fcf62fb9dc)

11. Next we need to check that the user input is valid or not , to do that we create a function called ValidateUserInput()
    
    - we need to provide user input as an argument in the function and we want it to return 3 boolean values.
    - to check if the name is valid or not ,we find out length of the strings firstname and lastname and check if they are greater than 1 , if they are then the name is valid.
    - to check if the NoOfTickets are valid or not we need to check if it is less than remaining tickets (remTicket) , if it is then the No. Of tickets are valid.
    - to check if the email is valid or not we need to check if it contains "@" and "." , if it does then email is considered valid .

      ```go
      func ValidateUserInput(firstname string , lastname string , NoOfTickets uint , email string , remTicket uint ) (bool , bool , bool) {
          isValidName := len(firstname)>1 && len(lastname)>1
          isValidTicket := NoOfTickets < remTicket
          isValidemail := strings.Contains(email , "@") && strings.Contains(email , ".") // to use strings package we need to declare import "strings" 

          return isValidName , isValidTicket , isValidemail
      }
      ```
12. now we have declared and initialized the function but let's declare other important functions first and then call them in our main function.
    But before that we need to let our program keep running for more than one iteration until all of our tickets are booked and we also have to find a way to store the data 
    that we recieved .

13. To keep our program running till all the tickets are booked we are going to use for loop that runs until the remaing tickets are greater than 0 .

    ```go
    for <condition> {

    }
    ``` 
    
    ![1007](https://github.com/SagarSingh2003/newME-git/assets/129133613/f09342c1-6202-4a09-9fa2-bcd54387803b)



14. To store the data we are using a struct that is going to store firstname, lastname ,email and No. of tickets 

     i. first we need to create the struct outside the main function named "userdata" using the type keyword and mention what type of data we want to store in what variable.

      ```go
      type  userdata struct{
          firstName string
          lastName  string
          NOOFTICKETS uint
          Email  string
      }
     ```
     
    ii. now outside the main function we're going to make an empty slice/list of type "userdata" using make function to store the user data, so that the data does not get 
        lost in the next iteration.

       ```go
       var booking = make([]userdata, 0) //here 0 is the length of the slice , but when we add values it will expand accordingly.
       ```
    iii. now inside the main function we need to store the data entered by the user in that iteration . For this we create a "userData" struct of type "userdata" that we 
         declared above

       ```go
       var userData = userdata{
           firstName: firstname,
           lastName: lastname,
           NOOFTICKETS: NoOfTickets,
           Email: email,
       }
       ```
       this is what our func mainlooks like right now:
    
       ![1008](https://github.com/SagarSingh2003/newME-git/assets/129133613/aa70762c-9d34-474e-be42-bf6ea002c955)

15. Next we're gonna create a function naming GetFirstNames that stores firstnames of the users in a slice.

    - we're later eventually going to store the data in booking slice that we had declared after declaring type "userdata" struct, so to extract the first names we are going 
      to use the range keyword to get all the structs stored in booking and extract each firstname from their respective struct.

      ```go
      func GetFirstNames() {

          Firstnames := []string{} //we created Firstnames empty slice

          for _, book := range booking {  
              Firstnames = append(Firstnames, book.firstName)
          }
          fmt.Println("The first names are :", Firstnames)
      }
      ```
      
16. The last function that we're gonna create will be SendTicekt() , as the name suggests this function will display a message in the terminal that tickets have been sent after the userinput is checked.

    - when a ticket is created and sent it takes some time ,so we are going to emulate it by giving a 10 second time gap 
      before printing the confirmation message. we're gonna use "Sleep" function of the "time" package to do this.
    - here we are also gonna use "Sprintf" function of fmt package to store the print message and print it later on . (Sprintf unlike printf does not print anyting on the   
      terminal but it allows the string to be formatted and stores the string so that it can be used later on)
    - we're also gonna use go routines later while calling this function so that the 10second sleep does not affect the functioning of the rest of the program.

    ```go
    func SendTicket(NoOfTickets uint , firstname string , lastname string , email string){
       var ticket = fmt.Sprintf("%v tickets for user %v %v at %v",NoOfTickets , firstname , lastname , email)
       time.Sleep(10*time.Second)
       fmt.Println("#############")
       fmt.Printf("Sending tickets : %v \n", ticket)
       fmt.Println("#############") 
    }
    ```
    
17. So let's call the functions one by one now and add functionality to our program .

    i. Let's begin by calling ValidateUserInput() which returns us three bool values isValidName , isValidTicket, isValidemail

    ![image](https://github.com/SagarSingh2003/newME-git/assets/129133613/3772135c-c22b-4e92-aa8a-cd30ccca58b5)

    ```go
    
    isValidName , isValidTicket , isValidemail := ValidateUserInput(firstname , lastname , NoOfTickets , email)

    ```

   ii. now we have the information of validity of user data stored in the three variables , by using if-else statements we will check if the values are valid or not if any 
       of the values are not valid then that iteration will end displaying a message that will inform what value is not valid and the loop will move to the next iteration.

   ```go
    if isValidName && isValidTicket && isValidemail { //the condition is fulfilled if each of them have value true
        fmt.Printf("Congratulations !! %d tickets have been sent to the user %v \n",NoOfTickets , firstname)
        booking = append(booking, userData) //if the values are valid then we store the values in the booking list
        remTicket -= NoOfTickets // we subtract the no. of tickets from remaining tickets to find the new remaining tickets

        SendTicket(NoOfTickets , firstname ,lastname , email)

        GetFirstNames()  
    }else {

        if !isValidName{  //we are using nested if to know what is the exact error to let the user know what is wrong
        fmt.Println("OOps !! The name is too short , Please try again .")
        }
        if !isValidTicket{
        fmt.Printf("OOps !! We only have %d tickets remaing , please try again ." , NoOfTickets)  
        }
        if !isValidemail{
        fmt.Println("Oops !! The email is invalid, Please enter a Valid email address .")
        }
        continue //this will break the loop and move to the next iteration
   }

``` 
iii. The program is almost completely functional but there is one problem , as we mentioned above while our program is running and when the SendTicket() function is called 
     the whole program stops for 10seconds and we have to wait until that time is over for the next iteration to run. 
     to solve this problem we are gonna use go routines that will allow us to run SendTicket function in a seperate thread that will execute seperately and won't interfere 
     with the rest of the program .
       
 - we need to declare a waitgroup function from sync package outside func main because we don't want our program to end before all the threads have completed their 
       execution
       
 - to learn more about waitgroups you can go [here](https://github.com/Yashsharma1911/Golang-tutorial/tree/main/WaitGroups)
       
 - to secify a function as goroutine we need to mention "go" before the function , that's how easy it is.

 - to declare a waitgroup outside func main

 ```go
  var Wg = sync.WaitGroup{}
 ```
    
 - to declare the go routine

 ```go
  go SendTicket(NoOfTickets, firstname, lastname, email)
 ```

 - we need to add  Add() , Wait() , Done() functions of waitgroup in their respective positions 

   i. Wg.Add(1) just before the go routine so that the counter of waitgroup increases by 1

    ![image](https://github.com/SagarSingh2003/newME-git/assets/129133613/49f77f6e-4288-4458-b0c2-c5838515910e)
    
   ii. Wg.Wait() just before the func main ends so that the program does not end until all threads are over

    ![image](https://github.com/SagarSingh2003/newME-git/assets/129133613/36905dfb-7ed5-4f6d-afd3-7e98c27d5358)

   iii. Wg.Done() just before the func SendTicket() ends so that the counter decreases by 1 and the compiler knows that the thread has completed execution.

   ![image](https://github.com/SagarSingh2003/newME-git/assets/129133613/cf058791-c331-4764-9701-7fcf96547f29) 


we've finally completed our program and now it's time to build it into an executable file , to do that : 


-  in the terminal write
    ```go
    go build . 
    ```
- to run the exe file just open the terminal in the directory and run the command

   ```go
   ./bookingapp4
   ```
- to directly run the main.go file without building it into an executable we need to write in the terminal

   ```go
   go run main.go
   ```

#### Thats it !! You've completed your ticket booking application try to give faulty user input and see what kind of output you get and have fun with it.
 
