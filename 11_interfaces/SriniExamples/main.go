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

//		e = Intern{"John"}
//		e.PrintName()
//	}
// package main

// import "fmt"

// // Define the interface Notifier with Notify Method
// type Notifier interface {
// 	Notify()
// }

// // Now define 2 Structs email and SMS
// type Email struct{}
// type SMS struct{}

// // Define the Notify functions on each Struct
// func (e Email) Notify() { fmt.Println("Sending Email Notification") }
// func (s SMS) Notify()   { fmt.Println("Sending SMS Notification") }

// // Function that accepts any Notifier type
// func sendNotification(n Notifier) {
// 	n.Notify()
// }

// // In Main use the sendNotification function for each of the Structs above
// func main() {
// 	sendNotification(Email{}) // Output: Sending Email Notification
// 	sendNotification(SMS{})   // Output: Sending SMS Notification
// }

package main

import "fmt"

// Define an interface
type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

// Real implementation
type Stripe struct{}

func (s Stripe) ProcessPayment(amount float64) {
	fmt.Println("Processing payment through Stripe:", amount)
}

// Mock implementation for testing
type MockPayment struct{}

func (m MockPayment) ProcessPayment(amount float64) {
	fmt.Println("Mock payment processed:", amount)
}

// Function that depends on the interface
func checkout(p PaymentProcessor, amount float64) {
	p.ProcessPayment(amount)
}

func main() {
	checkout(Stripe{}, 100.50)      // Output: Processing payment through Stripe: 100.5
	checkout(MockPayment{}, 100.50) // Output: Mock payment processed: 100.5
}
