package main

import "fmt"

func main() {

	var a = "string var declartion"
	fmt.Println(a)

	//multiple declartions i.e assigning 2 var with values
	var b, c int = 1, 2
	fmt.Println(b, c)

	// var declared as bool type
	var d = true
	fmt.Println(d)

	// var declared without a value (also known as zero-valued), by default value = 0
	var e int
	fmt.Println(e)

	// using ":=" shorthand way to declare variable
	//string
	f := "shorthand declartion"
	fmt.Println(f)

	//integer
	g := 67
	fmt.Println(g)

}
