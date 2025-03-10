package main

import "fmt"

var poInt int

func square(ptr *int) {
	*ptr = *ptr * *ptr
}

func main() {

	poInt = 16
	fmt.Println("initial poInt:", poInt)

	square(&poInt)

	fmt.Println("squared poInt:", poInt)
	fmt.Println("poInt pointer:", &poInt)
}
