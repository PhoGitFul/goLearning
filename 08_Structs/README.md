// Program to use function as a field  of struct

package main

import "fmt"

// initialize the function Cube
// This function reauires 3 integers input, and will output an integer, no logic for the function is 
// defined or required at this point we are just setting up the skeleton of the function here
type Cube func(int, int, int) int

// create structure called cubePara
// Last item in the Struct is a function of type Cube defined above.
type cubePara struct {
	length  int
	breadth int
	width   int
	color   string
	qube    Cube // Cube function as a field of struct
}

func main() {
	// assign values to struct variables
	result := cubePara{
		length:  10,
		breadth: 20,
		width:   3,
		color:   "Blue",
		qube: func(len int, bread int, wid int) int {
			return len * bread * wid
// it is in here that the function of type cube is defined and the logic added
		},
	}

	fmt.Println("Color of Cube: ", result.color)
	fmt.Println("Area of Cube: ", result.qube(result.length, result.breadth, result.width))
    // We are using the variables set within the result struct itself to pass to the result.qube struct 
    // function.
}

Exported fields
Now let's learn what is exported and unexported fields in a struct. Same as the rules for variables and functions, if a struct field is declared with a lower case identifier, it will not be exported and only be visible to the package it is defined in.

type Person struct {
	FirstName, LastName  string
	Age                  int
	zipCode              string
}
So, the zipCode field won't be exported. Also, the same goes for the Person struct, if we rename it as person, it won't be exported.

type person struct {
	FirstName, LastName  string
	Age                  int
	zipCode              string
}