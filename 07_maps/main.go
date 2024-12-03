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
// 	"b": 1, //Trailing commas are required on last entry
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

// func main() {
// 	var m = map[string]User{
// 		"a": {"Peter"},
// 		"b": {"Seth"},
// 	}

// 	m["c"] = User{"Steve"}

// 	fmt.Println(m)
// }

// func main() {
// 	var m = map[string]User{
// 		"a": {"Peter"},
// 		"b": {"Seth"},
// 	}

// 	m["c"] = User{"Steve"}
// 	c := m["c"]
// 	fmt.Println("Key c:", c)

// 	d := m["d"]
// 	fmt.Println("Key d:", d)
// }

// func main() {
// 	var m = map[string]User{
// 		"a": {"Peter"},
// 		"b": {"Seth"},
// 	}

// 	m["c"] = User{"Steve"}
// 	c, ok := m["c"]
// 	fmt.Println("Key c:", c, ok)
// 	d, ok := m["d"]
// 	fmt.Println("Key d:", d, ok)

// 	// m["a"] = "Roger"  Didn't like any of these 3 when trying to update the map
// 	// m["a"]: {"Roger"}
// 	// "a": {"Roger"}
// 	m["a"] = User{"Roger"}
// 	fmt.Println(m)

// 	delete(m, "b")
// 	fmt.Println(m)
// }

// func main() {
// 	var m = map[string]User{
// 		"a": {"Peter"},
// 		"b": {"Seth"},
// 	}
// 	m["c"] = User{"Steve"}
// 	for key, value := range m {
// 		fmt.Printf("Key: %s, Value: %v\n", key, value)
// 	}
// }

// func main() {
// 	var m1 = map[string]User{
// 		"a": {"Peter"},
// 		"b": {"Seth"},
// 	}
// 	m2 := m1
// 	m2["c"] = User{"Steve"}
// 	fmt.Println(m1)
// 	fmt.Println(m2)
// }

func main() {
	var m = map[string]User{
		"a": {"Peter"},
		"b": {"Seth"},
	}
	m["c"] = User{"Steve"}
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}
}
