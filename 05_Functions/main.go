package main

import "fmt"

func main() {
	fmt.Println("Add function to sum two numbers")
	i := add(17, 34)
	fmt.Println(i)
	fmt.Println("\nString function")
	q, r := hw("Hello", "World")
	fmt.Println(q)
	fmt.Println(r)

	fmt.Println("\nFunction as a value")
	funcValue()

	fmt.Println("\nAnonymous Function as a value")
	anonFunction()

	fmt.Println("\nReturn a value from function")
	johnPrice := computePrice(145.90, 3)
	paulPrice := computePrice(26.32, 10)
	robPrice := computePrice(189.21, 2)
	total := johnPrice + paulPrice + robPrice
	fmt.Printf("TOTAL : $%0.2f\n", total)

	fmt.Println("\nClosures")
	fmt.Println("Before assigning 'add' to closureFunc")
	add := closureFunc()
	fmt.Println("After assigning 'add' to closureFunc")
	add(5)
	add(2)
	fmt.Println(add(10))
	fmt.Println("\nEach call to the function uses the same variable within the function hence we can send diff values to the function and get the addition of the values sent")

	fmt.Println("\nBefore addition function call assigned")
	addition := closureFunc()
	fmt.Println("So as you can see above, the addition function call assigned to closureFunc means the variables within are reset")
	addition(1)
	addition(1)
	addition(1)
	fmt.Println(addition(1))

	fmt.Println("\nVaradic Function")
	sum := multiAdd(1, 2, 3, 5, 8, 13, 21)
	fmt.Println(sum)

}
func add(x int, y int) int {
	fmt.Println(x, "+", y, "=")
	return x + y
}

func hw(p1 string, p2 string) (string, string) {
	msg1 := fmt.Sprintf("%s function in the %s", p1, p2)
	msg2 := fmt.Sprintf("%s %s", p2, p1)
	return msg1, msg2
}

func funcValue() {
	thisfn := func() {
		fmt.Println("inside thisfn function")
	}

	thisfn()
}

func anonFunction() {
	func() {
		fmt.Println("inside anon function")
	}()
}

func computePrice(rate float32, nights int) float32 {
	return rate * float32(nights) //Need to convert int nights to same type as rate, which is float32 for calc to work
}

func closureFunc() func(int) int {
	sum := 0
	fmt.Println(sum)
	fmt.Println("So this is just the initialisation part which is only performed once. When the function is assigned.")

	return func(v int) int {
		sum += v
		fmt.Println(v)
		fmt.Println(sum)
		fmt.Println()
		return sum
	}
}

func multiAdd(values ...int) int {
	sum := 0
	for i, v := range values {
		sum += v
		fmt.Println(i, v)

	}
	return sum
}
