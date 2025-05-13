package main

import "fmt"

// Define an interface
type Product interface {
	GetPrice() float64
	GetDiscountPrice() float64
}

// Menswear Product
type MensWear struct {
	amount float64
}

// Womenswear Product
type WomensWear struct {
	amount float64
}

// Define the functions on each Struct
func (m MensWear) GetPrice() float64           { return m.amount }
func (m MensWear) GetDiscountPrice() float64   { return m.amount * (1 - 0.1) }
func (w WomensWear) GetPrice() float64         { return w.amount }
func (w WomensWear) GetDiscountPrice() float64 { return w.amount * (1 - 0.2) }

// Function to print original and discounted prices

func printPrices(p Product) {
	fmt.Printf("Original Price: £%0.2f, Discounted Price: £%0.2f\n", p.GetPrice(), p.GetDiscountPrice())
	fmt.Println()
}

func main() {
	mensJeans := MensWear{20}
	mensShirt := MensWear{15}

	fmt.Println("Mens Jeans -")
	printPrices(mensJeans)
	fmt.Println("Mens Shirt -")
	printPrices(mensShirt)

	womensBlouse := WomensWear{47}
	womensDress := WomensWear{63}

	fmt.Println("Womens Blouse -")
	printPrices(womensBlouse)
	fmt.Println("Womens Dress -")
	printPrices(womensDress)

}
