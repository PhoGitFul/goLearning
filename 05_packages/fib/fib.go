package fib

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Fibs(value int) {
	var a [51]int
	a[0] = 0
	a[1] = 1
	p := message.NewPrinter(language.English)
	if value < 0 {
		fmt.Println("Cannot handle negatives!")
		return
	}
	if value < 2 {
		fmt.Println("Fibonacci Number", value, "is :")
		p.Printf("%d\n", a[value])
		return
	}
	if value > 50 {
		fmt.Println("Cannot handle more than 50, so 50th fibonacci number is:")
		p.Printf("%d\n", 12586269025)
		return
	}
	for i := 2; i <= value; i++ {
		a[i] = a[i-2] + a[i-1]
	}
	fmt.Println("Fibonacci Number", value, "is :")
	p.Printf("%d\n", a[value])
}
