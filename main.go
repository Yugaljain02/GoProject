package main

import (
	"strings"
	"fmt"
)
func main(){

	conferenceName  := "Go Conference"
	const conferenceTickets int = 50
	var remainingtickets uint = 50
   
    bookings :=[]string{}

    fmt.Printf("conferenceTickets is %T,remainingTickets is %T, conferenceName %T\n",conferenceTickets,remainingtickets,conferenceName)
	fmt.Printf("welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available \n",conferenceTickets, remainingtickets)
    fmt.Printf("Get your tickets here to attend\n")
    
for {	
	var userName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter the first name: ")
	fmt.Scan(&userName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)
    
	remainingtickets = remainingtickets- userTickets
	 bookings = append(bookings, userName+" "+lastName)

	 fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve confirmation email at %v\n",userName,lastName,userTickets,email)
     fmt.Printf("%v tickets remaining for %v\n", remainingtickets,conferenceName)
    
	 firstName := []string{}
	 for _, booking := range bookings{

		var names = strings.Fields(booking)

		firstName = append(firstName,names[0])
	 }


	 fmt.Printf("this all are your bookings%v\n",bookings)
	}

}
