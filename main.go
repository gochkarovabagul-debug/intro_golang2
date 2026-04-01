package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferencename = "Go conference"
	const totalticketnumber = 50
	var remainingtickets = 50
	var name string
	var lastname string
	var orderedtickets int
	fmt.Println("welcome", conferencename, "booking application")
	fmt.Println("get your tickets here to attend")
	fmt.Printf("Dear %v %v user we have %v tickets and still have %v tickets\n", name, lastname, totalticketnumber, remainingtickets)
	var bookings = []string{}
	for {

		fmt.Println("enter your name:")
		fmt.Scan(&name)
		fmt.Println("enter your  last name:")
		fmt.Scan(&lastname)
		fmt.Println("enter your number of tickets:")
		fmt.Scan(&orderedtickets)
		remainingtickets = remainingtickets - orderedtickets

		bookings = append(bookings, name+" "+lastname)

		fmt.Printf("now we have %v number of tickets\n", remainingtickets)

		fmt.Printf("all our bookings :%v\n", bookings)

		var names = []string{}
		for _, bookings := range bookings {
			var name = strings.Fields(bookings)
			names = append(names, name[0])
		}
		fmt.Printf("The names of customers: %v\n", names)

	}
}
