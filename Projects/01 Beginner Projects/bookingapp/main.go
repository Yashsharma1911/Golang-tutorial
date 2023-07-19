package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var confname = "Gophers"
var totTickets uint = 50
var remTicket uint = totTickets // we used uint because remaining tickets
//can only have  positive values

type userdata struct {
	firstName   string
	lastName    string
	NOOFTICKETS uint
	Email       string
}

var booking = make([]userdata, 0) //here 0 is the length of the slice , but when we add values it will expand accordingly.

var Wg = sync.WaitGroup{}

func main() {

	for remTicket > 0 {
		GreetUsers()

		firstname, lastname, NoOfTickets, email := GetUserInput()

		var userData = userdata{
			firstName:   firstname,
			lastName:    lastname,
			NOOFTICKETS: NoOfTickets,
			Email:       email,
		}

		isValidName, isValidTicket, isValidemail := ValidateUserInput(firstname, lastname, NoOfTickets, email)

		if isValidName && isValidTicket && isValidemail { //the condition is fulfilled if each of them have value true
			fmt.Printf("Congratulations !! %d tickets have been sent to the user %v \n", NoOfTickets, firstname)
			booking = append(booking, userData) //if the values are valid then we store the values in the bookings list
			remTicket -= NoOfTickets            // we subtract the no. of tickets from remaining tickets to find the new remaining tickets

			Wg.Add(1)
			go SendTicket(NoOfTickets, firstname, lastname, email)

			GetFirstNames()
		} else {

			if !isValidName { //we are using nested if to know what is the exact error to let the user know what is wrong
				fmt.Println("OOps !! The name is too short , Please try again .")
			}
			if !isValidTicket {
				fmt.Printf("OOps !! We only have %d tickets remaing , please try again .", NoOfTickets)
			}
			if !isValidemail {
				fmt.Println("Oops !! The email is invalid, Please enter a Valid email address .")
			}
			continue //this will break the loop and move to the next iteration
		}

	}
	Wg.Wait()
}

func ValidateUserInput(firstname string, lastname string, NoOfTickets uint, email string) (bool, bool, bool) {
	isValidName := len(firstname) > 1 && len(lastname) > 1
	isValidTicket := NoOfTickets < remTicket
	isValidemail := strings.Contains(email, "@") && strings.Contains(email, ".")

	return isValidName, isValidTicket, isValidemail
}

func GreetUsers() {
	fmt.Printf("Welcome to %v ticket booking application!! \n", confname)
	fmt.Printf("Hurry! we only have %d seats remaining \n", remTicket)
}

// Func to take data from user
func GetUserInput() (string, string, uint, string) {
	var firstname string
	var lastname string
	var NoOfTickets uint
	var email string

	fmt.Println("Please Enter your firstname")
	fmt.Scan(&firstname)

	fmt.Println("Please Enter your lastname")
	fmt.Scan(&lastname)

	fmt.Println("Please Enter the No. of tickets you want to book")
	fmt.Scan(&NoOfTickets)

	fmt.Println("Please Enter your email")
	fmt.Scan(&email)

	return firstname, lastname, NoOfTickets, email
}

func GetFirstNames() {

	Firstnames := []string{} //we created Firstnames empty slice

	for _, book := range booking {
		Firstnames = append(Firstnames, book.firstName)
	}
	fmt.Println("The first names are :", Firstnames)
}

func SendTicket(NoOfTickets uint, firstname string, lastname string, email string) {
	var ticket = fmt.Sprintf("%v tickets for user %v %v at %v", NoOfTickets, firstname, lastname, email)
	time.Sleep(10 * time.Second)
	fmt.Println("#############")
	fmt.Printf("Sending tickets : %v \n", ticket)
	fmt.Println("#############")

	Wg.Done()
}
