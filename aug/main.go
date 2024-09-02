package main

import "fmt"

func main() {

	const myFirstName string = "Malcolm" // Long form way of declaring and setting a constant
	mySurname := "Graney"                //Short way to declare
	fmt.Println("My Name is", myFirstName, mySurname)

	numKids := 5 //Adding comment as I am declaring an int variable here
	fmt.Println("I have", numKids, "children")
	kidsName(numKids)

	var married = true
	var happy = true

	if married && happy {
		fmt.Println("I am happily married")
	} else {
		if !married {
			switch happy {
			case true:
				fmt.Println("I'm loving the single life")
			case false:
				fmt.Println("I'm single and depressed")
			default: //put just for fun as being boolean should never hit here
				fmt.Println("How did I get here? Must be a bug")
			}
		} else {
			fmt.Println("I want a divorce NOW!")
		}
	}
}

func kidsName(kid int) {
	name := [5]string{"Milly", "Jack", "John", "Apple", "Alex"}
	order := [5]string{"First", "Second", "Third", "Fourth", "Fifth"}
	for i := 0; i < kid; i++ { //this i identifier is local to the for loop only
		fmt.Println("My", order[i], "child is called", name[i])
	}
}
