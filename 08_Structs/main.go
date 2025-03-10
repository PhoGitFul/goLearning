// package main

// import "fmt"

// type Person struct {
// 	name string
// 	age  int
// }

// func main() {
// 	var pers1 Person
// 	var pers2 Person

// 	// Pers1 specification
// 	pers1.name = "Hege"
// 	pers1.age = 45

// 	// Pers2 specification
// 	pers2.name = "Cecilie"
// 	pers2.age = 24

// 	// Access and print Pers1 info
// 	fmt.Println("Name: ", pers1.name)
// 	fmt.Println("Age: ", pers1.age)

// 	// Access and print Pers2 info
// 	fmt.Println("Name: ", pers2.name)
// 	fmt.Println("Age: ", pers2.age)
// }
// Output
// Name:  Hege
// Age:  45
// Name:  Cecilie
// Age:  24

// package main

// import "fmt"

// type Person struct {
// 	name string
// 	age  int
// }

// func main() {
// 	var pers1 Person
// 	var pers2 Person

// 	// Pers1 specification
// 	pers1.name = "Hege"
// 	pers1.age = 45

// 	// Pers2 specification
// 	pers2.name = "Cecilie"
// 	pers2.age = 24

// 	// Print Pers1 info by calling a function
// 	printPerson(pers1)

// 	// Print Pers2 info by calling a function
// 	printPerson(pers2)
// }

// func printPerson(pers Person) {
// 	fmt.Println("Name: ", pers.name)
// 	fmt.Println("Age: ", pers.age)
// }
// Output
// Name:  Hege
// Age:  45
// Name:  Cecilie
// Age:  24

// package main

// import "fmt"

// // declare a struct
// type Person struct {
// 	name string
// 	age  int
// }

// func main() {

// 	// assign value to struct while creating an instance
// 	person1 := Person{"John", 25}
// 	fmt.Println(person1)

// 	// // define an instance
// 	// var person2 Person

// 	// // assign value to struct variables
// 	// person2 = Person{name: "Sara", age:  29}

// 	// The above way of assigning the value to the struct worked but a warning was shown:
// 	// "should merge variable declaration with assignment on next line (S1021)"
// 	// So when declaration and assign is done on same line it would look like this
// 	person2 := Person{name: "Sara", age: 29}

// 	// assign value to struct variables
// 	//person2 = Person{name: "Sara", age: 29}

// 	fmt.Println(person2)
// }
// Output
// {John 25}
// {Sara 29}

// Program to access the individual elements of struct

// package main

// import "fmt"

// func main() {
// 	// declare a struct
// 	type Rectangle struct {
// 		length  int
// 		breadth int
// 		width   int
// 	}

// 	// declare instance rect1 and defining the struct
// 	rect := Rectangle{22, 12, 2}

// 	// access the length of the struct
// 	fmt.Println("Length:", rect.length)

// 	// access the breadth of the struct
// 	fmt.Println("Breadth:", rect.breadth)

// 	// access the width of the struct
// 	fmt.Println("Width:", rect.width)

// 	area := rect.length * rect.breadth
// 	fmt.Println("Area:", area)

// 	cubeArea := rect.length * rect.breadth * rect.width
// 	fmt.Println("Cube Area:", cubeArea)
// }
// Output =
// Length: 22
// Breadth: 12
// Width: 2
// Area: 264
// Cube Area: 528

// Program to use function as a field  of struct

// package main

// import "fmt"

// // initialize the function Cube,
// type Cube func(int, int, int) int

// // create structure
// type cubePara struct {
// 	length  int
// 	breadth int
// 	width   int
// 	color   string
// 	qube    Cube // Cube function as a field of struct
// }

// func main() {
// 	// assign values to struct variables
// 	result := cubePara{
// 		length:  10,
// 		breadth: 20,
// 		width:   3,
// 		color:   "Blue",
// 		qube: func(len int, bread int, wid int) int {
// 			return len * bread * wid
// 		},
// 	}

// 	fmt.Println("Color of Cube: ", result.color)
// 	fmt.Println("Area of Cube: ", result.qube(result.length, result.breadth, result.width))
// }
// Output = 	Color of Cube:  Blue
// 			Area of Cube:  600

// Embedding
// package main

// import "fmt"

// type Person struct {
// 	FirstName, LastName string
// 	Age                 int
// }

// type SuperHero struct {
// 	Person
// 	Alias string
// }

// func main() {
// 	s := SuperHero{}
// 	s.FirstName = "Bruce"
// 	s.LastName = "Wayne"
// 	s.Age = 40
// 	s.Alias = "batman"
// 	fmt.Println(s)
// }
//Output = {{Bruce Wayne 40} batman}

// Composition

// package main

// import "fmt"

// type Person struct {
// 	FirstName, LastName string
// 	Age                 int
// }
// type SuperHero struct {
// 	Pers  Person
// 	Alias string
// }

// func main() {
// 	p := Person{"Bruce", "Wayne", 40}
// 	s := SuperHero{p, "batman"}

// 	fmt.Println(s)
// }

// Output = {{Bruce Wayne 40} batman}

// Srini Examples
// nested structs

// package main

// import "fmt"

// type details struct {
// 	number int
// }
// type employee struct {
// 	id   int
// 	name string
// 	details
// }

// func main() {
// 	var e1 employee
// 	fmt.Printf("\n e1 is %v", e1)

// 	e2 := employee{name: "john", id: 12345, details: details{number: 2}}
// 	fmt.Printf("\n e2 is %v", e2)

// 	e3 := employee{id: 4321, name: "david", details: details{number: 9}}
// 	fmt.Printf("\n e3 is %v", e3)

// 	e4 := employee{72, "ozzie", details{number: 7}}
// 	fmt.Printf("\n e4 is %v", e4)
// 	//If we don't specify the type when setting the Struct then we need to put the fields in the same order as they are defined

// 	type addr struct {
// 		houseNo  int
// 		postCode string
// 		city     string
// 	}
// 	type worker struct {
// 		id      int
// 		name    string
// 		address addr
// 	}
// 	var w1 worker
// 	fmt.Printf("\n w1 is %v", w1)

// 	w2 := worker{name: "Ringo", id: 4, address: addr{houseNo: 32, postCode: "W4 3AP", city: "Ealing"}}
// 	fmt.Printf("\n w2 is %v", w2)

// 	fmt.Println("worker id ", w2.id)
// 	fmt.Println("worker name ", w2.name)
// 	fmt.Println("worker house no ", w2.address.houseNo)
// 	fmt.Println("worker postcode ", w2.address.postCode)
// 	fmt.Println("worker city ", w2.address.city)

// Output =
// e1 is {0  {0}}
//	e2 is {12345 john {2}}
//	e3 is {4321 david {9}}
//	e4 is {72 ozzie {7}}
//	w1 is {0  {0  }}
//	w2 is {4 Ringo {32 W4 3AP Ealing}}worker id  4
// worker name  Ringo
// worker house no  32
// worker postcode  W4 3AP
// worker city  Ealing

//Methods - Functions referenced with a particular struct are methods of that struct

package main

import "fmt"

type employee struct {
	id   int
	name string
}

func (e employee) PrintDetails() {
	fmt.Printf("\n employed id %d and name %s", e.id, e.name)
}
func main() {
	var e1 employee
	e1.PrintDetails()
	e2 := employee{id: 1234, name: "john"}
	e2.PrintDetails()
	e3 := employee{4321, "david"}
	e3.PrintDetails()
	e4 := employee{}
	e4.PrintDetails()

	// Output =
	// employed id 0 and name
	// employed id 1234 and name john
	// employed id 4321 and name david
	// employed id 0 and name %
}
