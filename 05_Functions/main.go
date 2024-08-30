package main

import "fmt"

func main() {
	i := add(17, 34)
	fmt.Println(i)
}

// return sum of a and y
func add(x int, y int) int {
	return x + y
}

//func main() {
//	s, i := myFunction("Hello")
//	fmt.Println(s, i)
//}
//func myFunction(p1 string) (string, int) {
//	msg := fmt.Sprintf("%s function", p1)
//	return msg, 10
//}
