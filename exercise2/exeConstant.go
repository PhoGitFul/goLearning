package main

import (
	"fmt"
)

func main() {
	const hotelName string = "Gopher Hotel" //typed constant
	const longitude = 24.806078             //untyped constant default type to float
	const latitude = -78.243027             //untyped constant default type to float
	var occupancy int = 12                  //typed variable

	fmt.Println(hotelName, longitude, latitude, occupancy)
	//fmt.Println(reflect.TypeOf(longitude))
	//fmt.Println(reflect.TypeOf(latitude))
	//	fmt.Println("Constant Exercise")
	//	fmt.Println("Hotel:", hotelName, "is located at longitude -", longitude, ",", "latitude -", latitude, ",", "with occupancy of ", occupancy)

}
