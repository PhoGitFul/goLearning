package main

import "fmt"

// import (
// 	"fmt"
// )

// func main() {
// 	x := 11

// 	if x > 10 {
// 		fmt.Println("x =", x, "and so is gt 10")
// 	} else if x > 5 {
// 		fmt.Println("x =", x, "and so is gt 5")
// 	} else {
// 		fmt.Println("x =", x, "and falls into the else and so must be lt 5")
// 	}
// }
// import (
// 	"fmt"
// 	"math/rand"
// )

// func main() {
// 	//	rand.Seed(time.Now().UTC().UnixNano())
// 	var agePaul, ageJohn int = rand.Intn(110), rand.Intn(110)
// 	// const ageJohn, agePaul int = 32, 32
// 	fmt.Println("John is", ageJohn, "years old")
// 	fmt.Println("Paul is", agePaul, "years old")
// 	if agePaul > ageJohn {
// 		fmt.Println("Paul is older than John")
// 	} else if agePaul == ageJohn {
// 		fmt.Println("John and Paul are the same age")
// 	} else {
// 		fmt.Println("Paul is younger than John")
// 	}
// }
// if else can be split out as below in Shitul's example
// func main() {
//     var agePaul, ageJohn int = rand.Intn(110), rand.Intn(110)
//     fmt.Println("John is", ageJohn, "years old")
//     fmt.Println("Paul is", agePaul, "years old")
//     if agePaul > ageJohn { //code 1
//         fmt.Println("Paul is older than John")
//     } else { // code 2
//         // another if/else structure
//         if agePaul == ageJohn {   // code 3
//             fmt.Println("Paul and John have the same age")
//         } else {
//             fmt.Println("Paul is younger than John")
//         }
//     }
// }

// func main() {
// 	day := "tuesday"

// 	switch day {
// 	case "monday":
// 		fmt.Println("time to work!")
// 	case "friday":
// 		fmt.Println("let's party")
// 	default:
// 		fmt.Println("browse memes")
// 	}
// }

func main() {
	i := 2 //This identifier is not used within the for loop hence why the value it is set to is still
	//equal to 2 even after the i identifier is used in teh for loop and increased from 0 to 9
	for i := 0; i < 10; i++ { //this i identifier is local to the for loop only
		fmt.Println(i)
	}
	fmt.Println(i)
}
