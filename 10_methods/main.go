// package main

// import "fmt"

// type employee struct {
// 	id   int
// 	name string
// }

// func (e employee) PrintDetails() {
// 	fmt.Printf("\n employee id %d and name %s", e.id, e.name)
// }
// func main() {
// 	var e1 employee
// 	e1.PrintDetails()
// 	e2 := employee{id: 1234, name: "john"}
// 	e2.PrintDetails()
// 	e3 := employee{4321, "david"}
// 	e3.PrintDetails()
// 	e4 := employee{}
// 	e4.PrintDetails()
// }

// Results
//mal@ip-10-192-241-246 10_methods % go run main.go
//
//  employee id 0 and name
//  employee id 1234 and name john
//  employee id 4321 and name david
//  employee id 0 and name %
//

// --------------------
// Encapsulation example
// Use methods to restrict direct access to fields (getter/setter).
// --------------------
// package main

// import (
// 	"fmt"
// 	//	"Models"
// 	"goLearning/10_methods/Models"
// 	// "project/Models"
// )

//	func main() {
//		e1 := &Models.Employee{}
//		e1.SetID(123)
//		fmt.Println(e1.GetID())
//	}
//
// Results when running the Models methods
// mal@ip-10-192-241-246 10_methods % go run main.go
// 123
//
// --------------------
// Builders or Chaining example
// Use method chaining for clean object creation.
// --------------------
// package main

// import "fmt"

// type Rectangle struct {
// 	width, height float64
// }

// // Builder method to set width
// func (r *Rectangle) SetWidth(w float64) *Rectangle {
// 	r.width = w
// 	return r // Returning pointer for method chaining
// }

// // Builder method to set height
// func (r *Rectangle) SetHeight(h float64) *Rectangle {
// 	r.height = h
// 	return r
// }

// // Method to calculate Area
// func (r Rectangle) Area() float64 {
// 	return r.width * r.height
// }

// // Display the rectangle details
// func (r Rectangle) Show() {
// 	fmt.Println("Rectangle - Width:", r.width, "Height:", r.height, "Area:", r.Area())
// }

// func main() {
// 	rect := &Rectangle{} // Empty rectangle instance

// 	// Using builder pattern
// 	rect.SetWidth(2).SetHeight(4.8).Show()
// 	rect.SetWidth(5).SetHeight(10).Show()
// 	rect.SetHeight(2).Show()
// }

// Output when using Pointers e.g. func (r *Rectangle) SetWidth(w float64) *Rectangle {
// Rectangle - Width: 2 Height: 4.8 Area: 9.6
// Rectangle - Width: 5 Height: 10 Area: 50
// Rectangle - Width: 5 Height: 2 Area: 10
// Output when NOT using Pointers e.g. func (r Rectangle) SetWidth(w float64) Rectangle {
// Rectangle - Width: 2 Height: 4.8 Area: 9.6
// Rectangle - Width: 5 Height: 10 Area: 50
// Rectangle - Width: 0 Height: 2 Area: 0
// When pointers aren't used we can see a new instance of rect is created and as such has not got
// Width set on the last call, whereas when we use a Pointer we are referencing the same instance
// and so Width is still set as 5 from previous call.
//
// --------------------
// Behaviour Implementation example
// Attach meaningful actions to types.
// --------------------
package main

import "fmt"

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	rect := Rectangle{5, 10}
	fmt.Println("Area:", rect.Area()) // Output: Area: 50
}
