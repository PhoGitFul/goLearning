// package main

// import "fmt"

// func main() {
// 	var p *int
// 	fmt.Println(p)
// }
// Output is nil as we have not set the pointer to a value yet.
// $ go run main.go
// nil

// package main

// import "fmt"

// func main() {
// 	a := 10

// 	var p *int = &a

// 	fmt.Println("address:", p)
// }
// Output
// $ go run main.go
// address: 0x1400000e0b0

// Dereferencing
// package main

// import "fmt"

// func main() {
// 	a := 10

// 	var p *int = &a

// 	fmt.Println("address:", p)
// 	fmt.Println("value:", *p)
// }
// Output
// $ go run main.go
// address: 0x1400000e0b0
// value: 10

// Change the value using the pointer
// package main

// import "fmt"

// func main() {
// 	a := 10

// 	var p *int = &a

// 	fmt.Println("before", a)
// 	fmt.Println("address:", p)

// 	*p = 20
// 	fmt.Println("after:", a)
// }

// Output
// $ go run main.go
// address: 0x1400000e0b0
// value: 10
// after: 20

// New function to initialise a pointer
package main

import "fmt"

func main() {
	p := new(int)
	*p = 100

	fmt.Println("value", *p)
	fmt.Println("address", p)
}

// Output
// $ go run main.go
// value 100
// address 0x14000112018
