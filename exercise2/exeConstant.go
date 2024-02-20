package main

import (
	"fmt"
)

func main() {
	const hotelName string = "Gopher Hotel" //typed constant
	const longitude = 24.806078             //untyped constant default type to float
	const latitude = -78.243027             //untyped constant default type to float
	var occupancy int = 12                  //typed variable

	fmt.Println("Hotel:     ", hotelName)
	fmt.Println("Longitude: ", longitude)
	fmt.Println("Latitude:  ", latitude)
	fmt.Println("Occupancy: ", occupancy)
}
