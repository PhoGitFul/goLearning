package main

import "fmt"

//func main() {

// var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
// b := a

// fmt.Println(a)
// fmt.Println(b)
//}

//func main() {

// b["year"] = "1970"
// fmt.Println("After change to b:")

// fmt.Println(a)
// fmt.Println(b)

//func main() {
// var m = make(map[string]int)
// fmt.Println(m)
//}

//func main() {
// var m = map[string]int{
// 	"a": 0,
// 	"b": 1,
// }
// fmt.Println(m)
//}

type User struct {
	Name string
}

// func main() {
// 	var m = map[string]User{
// 		"a": User{"Peter"},
// 		"b": User{"Seth"},
// 	}
// 	fmt.Println(m)
// }

func main() {
	var m = map[string]User{
		"a": {"Peter"},
		"b": {"Seth"},
	}

	m["c"] = User{"Steve"}

	fmt.Println(m)
}
