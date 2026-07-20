package main

import "fmt"

func main(){
	var conferenceName  string = "Go Conference"
	const conferenceTickets int = 50
	var remainingtickets uint = 50
	
    fmt.Printf("conferenceTickets is %T,remainingTickets is %T, conferenceName %T\n",conferenceTickets,remainingtickets,conferenceName)
	fmt.Printf("welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available \n",conferenceTickets, remainingtickets)
    fmt.Printf("Get your tickets here to attend")

	var userName string
	 var userTickets int

	userName = "Tom"
	 userTickets =2
	 fmt.Printf("User %v booked %v tickets \n", userName, userTickets)
}


