// package main

// type Employee interface {
// 	PrintName()
// }

// type Intern struct {
// 	Name string
// }

// func (i Intern) PrintName() {
// 	println(i.Name)
// }

// func main() {
// 	var e Employee

// 	e = Intern{"John"}
// 	e.PrintName()
// }

// Example taken and played with from DigitalOcean
// Please excuse the number of comments I need these for my reference when looking back at what I have learned
// In real functional Code would not include so many comments :-)
package main

import "fmt"

// This example defines an interface called Submersible that expects types having a Dive() method.
type Submersible interface {
	Dive()
}

// We then define a Shark type with a Name field and an isUnderwater boolean
// to keep track of the state of the Shark.
// Shark has 2 methods String() and Dive()

type Shark struct {
	Name string

	isUnderwater bool
}

// We also defined the String() method of the value receiver so that it could cleanly print
// the state of the Shark using fmt.Println by using the fmt.Stringer interface
// accepted by fmt.Println
func (s Shark) String() string {
	if s.isUnderwater {
		return fmt.Sprintf("%s is underwater", s.Name)
	}
	return fmt.Sprintf("%s is on the surface", s.Name)
}

// We define a Dive() method on the pointer receiver to Shark which
// modifies isUnderwater to true.
func (s *Shark) Dive() {
	s.isUnderwater = true
}

// We also used a function submerge that takes a Submersible parameter.
func submerge(s Submersible) {
	s.Dive()
}

//
// We could have the submerge function that takes a specific type and so pass the pointer to that type
// func submerge(s *Shark) {
// 	s.Dive()
// }

// Using the Submersible interface rather than a *Shark allows the submerge function
// to depend only on the behavior provided by a type.
// This makes the submerge function more reusable because you wouldn’t have to write new
// submerge functions for a Submarine, a Whale, or any other future aquatic inhabitants
// we haven’t thought of yet.
// As long as they define a Dive() method, they can be used with the submerge function

// Define a new type

type DeepSeaDiver struct {
	Name       string
	diverLevel string
}

// We also defined the String() method of the value receiver so that it could cleanly print
// the state of the Shark using fmt.Println by using the fmt.Stringer interface
// accepted by fmt.Println
func (d DeepSeaDiver) String() string {
	if d.diverLevel == "Zero" {
		return fmt.Sprintf("%s is not a Diver", d.Name)
	}
	return fmt.Sprintf("%s is a Level %s Diver", d.Name, d.diverLevel)
}

// We define a Dive() method on the pointer receiver to Shark which
// modifies isUnderwater to true.
func (d *DeepSeaDiver) Dive() {
	d.diverLevel = "One"
}

func detectSeaCreature(s Submersible) {
	if d, ok := s.(*DeepSeaDiver); ok {
		fmt.Printf("The Diver is called %s\n", d.Name)
	} else {
		fmt.Printf("He must be a Shark\n")
	}
}

func main() {
	// We define a variable s that is a pointer to a Shark type and immediately print s with fmt.Println
	s := &Shark{
		Name: "Sammy",
	}

	fmt.Println(s)
	// We pass s (the Shark type named "Sammy") to submerge function
	// which uses the Dive method that sets the boolean isUnderwater to true
	submerge(s)
	// and then called fmt.Println again with s as its argument to see the second part of the output printed, Sammy is underwater

	fmt.Println(s)
	detectSeaCreature(s)
	d := &DeepSeaDiver{
		Name:       "Dave",
		diverLevel: "Zero",
	}

	fmt.Println(d)
	// We pass s (the Shark type named "Sammy") to submerge function
	// which uses the Dive method that sets the boolean isUnderwater to true
	submerge(d)
	// and then called fmt.Println again with s as its argument to see the second part of the output printed, Sammy is underwater

	fmt.Println(d)
	detectSeaCreature(d)
}
