package main

import (
	"fmt"
	"time"
    "sync"
)
var conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = 50    
var bookings = make([]UserData, 0)

type UserData struct{
    firstName string
    lastName string
    email string
    numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main (){
    
    greetUsers()
        firstName, lastName, email, userTickets := getUserInput()

        isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)


        if isValidName && isValidEmail && isValidTicketNumber {
            bookTicket( userTickets , firstName , lastName , email)

            wg.Add(1)
            go sendTicket( userTickets , firstName , lastName , email)
            
            firstNames := getFirstNames()
            fmt.Printf("These are first names of all our bookings: %v \n", firstNames)


            if remainingTickets == 0 {
                fmt.Printf("Conference %v is booked out", conferenceName)
                // break
            }
        }else {
            if !isValidName{
                fmt.Println("First name or last name entered is too short")
            }
            if !isValidEmail{
                fmt.Println("Email address does not contain @ sign")
            }
            if !isValidTicketNumber{
                fmt.Println("Number of tickets is invalid.")
            }
        } 
        wg.Wait()
}

func greetUsers(){
    fmt.Printf("Welcome to %v booking application. \n", conferenceName)
    fmt.Printf("We have a total of %v tickets and %v are still available. \n", conferenceTickets, remainingTickets)
    fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string{
    firstNames := []string{}
    for _, booking := range bookings {
    firstNames = append(firstNames, booking.firstName)
    }
    return firstNames
    
}


func getUserInput() (string, string, string, uint){
    var firstName string 
    var lastName string 
    var email string 
    var userTickets uint

    fmt.Println("Enter your first name:")
    fmt.Scan(&firstName)

    fmt.Println("Enter your last name:")
    fmt.Scan(&lastName)

    fmt.Println("Enter your email:")
    fmt.Scan(&email)

    fmt.Println("Enter your number of tickets:")
    fmt.Scan(&userTickets)

    return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint,firstName string, lastName string, email string){
    remainingTickets = remainingTickets - userTickets

    var userData = UserData{
        firstName: firstName,
        lastName: lastName,
        email: email,
        numberOfTickets: userTickets,
    }

    bookings = append(bookings, userData)
    fmt.Printf("List of bookings %v \n", bookings)

    fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
    fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}

func sendTicket( userTickets uint, firstName string, lastName string, email string){
    time.Sleep(10* time.Second)
    var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
    fmt.Println("########")
    fmt.Printf("Sending ticket: to %v to email address %v \n", ticket, email)
    fmt.Println("########")
    wg.Done()
}