package main

import "fmt"

func main() {

	const myFirstName string = "Malcolm" // Long form way of declaring and setting a constant
	mySurname := "Graney"                // Short way to declare a var
	fmt.Println("My Name is", myFirstName, mySurname)

	numKids := 10 //Declaring an int variable here in short format
	if numKids > 0 && numKids < 5 {
		fmt.Println("I have", numKids, "children")
	}
	kidsName(numKids) // kidsName function to display the names of my kids, passing how many kids I have as the parameter
	// I could have moved the call to kidsName within the if numKids > 0 code and then not required the switch case witihin the function
	// but left it as is for now.
	var married bool = true //setting 2 boolean variables this one is long hand
	happy := true           //this one shorthand

	if married && happy { //If statment - if both booleans are true using the AND Logical Operator
		fmt.Println("I am happily married")
	} else { // Else fall into this section
		if !married { //Check if not married
			switch happy { //switch case for the 2 possible combinations when not married
			case true:
				fmt.Println("I'm loving the single life")
			case false:
				fmt.Println("I'm single and depressed")
			default: //put just for fun as being boolean should never hit here
				fmt.Println("How did I get here? Must be a bug")
			}
		} else { // Else last combo left means I must be married and not happy
			fmt.Println("I want a divorce NOW!")
		}
	}
}

func kidsName(kid int) { //there is 1 paramter expected to be passed in "kid" as int but nothing passed back out to calling app
	if kid < 1 || kid > 5 { //Using logical operator OR to catch values which would make the code fail
		switch {
		case kid < 1:
			fmt.Println("I have no kids")
		case kid > 5:
			fmt.Println("I have far too many kids -", kid, "!")
		}
	} else {
		name := [5]string{"Milly", "Jack", "John", "Apple", "Alex"}       // declare an array of 5 names
		order := [5]string{"First", "Second", "Third", "Fourth", "Fifth"} // declare an array of 5 positions
		for i := 0; i < kid; i++ {                                        // this i identifier is local to the for loop only
			fmt.Println("My", order[i], "child is called", name[i])
			// Given the number of kids passed in we loop round until i is less than the number of "kid"s variable.
			//Printing out the name and order from the 2 arrays
			//As i starts at 0 then we can be sure (as long as kid is > 0), that we have at least one pass through.
			//Improvement can be made to print a sentence stating I have no kids maybe if kid = 0, and also a default sentence printed
			//if kid sent is greater than 5.
		}
	}
}
