package main

import "fmt"

func main() {
	myMap := map[string]int{"Adam": 10, "Brian": 20, "Cathy": 30, "Des": 40, "Ed": 50}
	fmt.Println(myMap)

	myMap["Cathy"] = 25
	fmt.Println(myMap)

	m, ok := myMap["Ed"]
	fmt.Println("Key Ed:", m, ok)

	j, ok := myMap["Fred"]
	fmt.Println("Key Fred:", j, ok)
}
