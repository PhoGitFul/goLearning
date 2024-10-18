package main

import "fmt"

/*
	func main() {
		i := add(17, 34)
		fmt.Println(i)
	}

// return sum of a and y

	func add(x int, y int) int {
		return x + y
	}
*/
/*
func main() {
	q, r := myFunction("Hello", "World")
	fmt.Println(q)
	fmt.Println(r)
}
func myFunction(p1 string, p2 string) (string, string) { //%s is the holder for each string parameter in turn
	msg1 := fmt.Sprintf("%s function in the %s", p1, p2) // so in this example prints "Hello function in the World"
	msg2 := fmt.Sprintf("%s %s", p2, p1)                 // and this one prints "World Hello"
	return msg1, msg2
}

func main() {
	q, r := myFunction("Hello", "World")
	fmt.Println(q)
	fmt.Println(r)
}
func myFunction(p1 string, p2 string) (a string, b string) { // Named the returned output strings a & b
	a = fmt.Sprintf("%s function in the %s", p1, p2) // a and b returned to calling code and stored into q and r respectively
	b = fmt.Sprintf("%s %s", p2, p1)
	return
}
//
func main() {
	myFunction()
}
func myFunction() {
	fn := func() {
		fmt.Println("inside fn")
	}

	fn()
}
//
func main() {
	myFunction()
}

func myFunction() {
	func() {
		fmt.Println("inside anon fn")
	}()
}
//
func main() { // made fn from above an anonymous function, need the empty parantheses at end to call the function
	func() {
		fmt.Println("inside anonymous fn")
	}()
}
*/
/*
func main() {
	johnPrice := computePrice(145.90, 3) //*\label{funcEx1Us1}
	paulPrice := computePrice(26.32, 10) //*\label{funcEx1Us2}
	robPrice := computePrice(189.21, 2)  //*\label{funcEx1Us3}
	total := johnPrice + paulPrice + robPrice
	fmt.Printf("TOTAL : $%0.2f\n", total) // %0.2f formats the 'total' to 2 decimal places, \n terminates the print line and stops % character from being displayed in Terminal Output
}
func computePrice(rate float32, nights int) float32 { //*\label{funcEx1Dec}
	return rate * float32(nights) //Need to convert int nights to same type as rate, which is float32 for calc to work
}
*/
/*
func main() {
	fmt.Println("Before assigning add to myFunction")
	add := myFunction()
	fmt.Println("After assigning add to myFunction")
	add(5)
	add(2)
	fmt.Println(add(10))

	fmt.Println("Before addition function call assigned")
	addition := myFunction()
	fmt.Println("So as you can see above, the addition function call assigned to myFunction means the variables within are reset")
	addition(1)
	addition(1)
	addition(1)
}
func myFunction() func(int) int { // The anonymous function func(v int) is returned
	sum := 0 //This initialisation section is only performed once when called from an assigned function such as the add() or addition() above
	fmt.Println(sum)
	fmt.Println("So this is just the initialisation part which is only performed once. When the function is assigned.")
	//So v is the parameter passed from the assigned function, in this case from add() and addition().
	return func(v int) int {
		sum += v //+= Adds the value on the right to the value in the variable on the left and assigns it to the variable 'sum = sum +v'
		fmt.Println(v)
		fmt.Println(sum)
		fmt.Println()
		return sum // So this 'return sum' returns the calculation of the anonymous func
	}
}
*/
//
func main() {
	sum := add(1, 2, 3, 5, 8, 13, 21)
	fmt.Println(sum)
}
func add(values ...int) int {
	sum := 0
	for i, v := range values {
		sum += v
		fmt.Println(i, v)
		/*
		   for _, v := range values {  // do not need to use a counter such as i above if we are not going to use it anywhere so can just have
		   	sum += v                // underscore used
		*/
	}
	return sum
}
