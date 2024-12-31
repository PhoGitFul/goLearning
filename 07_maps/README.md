What is Maps?
In Golang, a map is a built-in data structure that stores a collection of key-value pairs. It is a hash table that allows you to store and retrieve data using a unique key. Maps are also known as hash maps, dictionaries, or associative arrays in other programming languages.
Here is an example of how to declare and initialize a map:



var m map[string]int
m = make(map[string]int)
Key Characteristics of Maps:
Maps are mutable, meaning they can be modified after creation.
Maps are unordered, meaning the order of key-value pairs is not guaranteed.
Maps are reference types, meaning they are passed by reference, not by value.
Maps are safe for concurrent use by multiple goroutines. 
Declaring and Initializing Maps:
Maps can be declared and initialized in several ways:
Using the make function: myMap := make(map[string]int)
Using a map literal: myMap := map[string]int{"a": 1, "b": 2}
Using a nil map: var myMap map[string]int (must be initialized before use)
Working with Maps
You can set key-value pairs using the typical name[key] = val syntax:



m["k1"] = 7
m["k2"] = 13
Printing a map with e.g. fmt.Println will show all of its key-value pairs:



fmt.Println("map:", m)
Get a value for a key with name[key]:



v1 := m["k1"]
fmt.Println("v1:", v1)
If the key doesn’t exist, the zero value of the value type is returned:



v3 := m["k3"]
fmt.Println("v3:", v3)
Maps Are References
Maps are references to hash tables.
If two map variables refer to the same hash table, changing the content of one variable affects the content of the other.
Example



var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
b := a
fmt.Println(a)
fmt.Println(b)
b["year"] = "1970"
fmt.Println("After change to b:")
fmt.Println(a)
fmt.Println(b)
Initialization
There are multiple ways to initialize a map.
make function
We can use the built-in make function, which allocates memory for referenced data types and initializes their underlying data structures.



func main() {
	var m = make(map[string]int)
	fmt.Println(m)
}



$ go run main.go
map[]
map literal
Another way is using a map literal.



func main() {
	var m = map[string]int{
		"a": 0,
    "b": 1,
	}
	fmt.Println(m)
}
Note that the trailing comma is required.



$ go run main.go
map[a:0 b:1]
As always, we can use our custom types as well.



type User struct {
	Name string
}
func main() {
	var m = map[string]User{
		"a": User{"Peter"},
		"b": User{"Seth"},
	}
	fmt.Println(m)
}
We can even remove the value type and Go will figure it out!



var m = map[string]User{
	"a": {"Peter"},
	"b": {"Seth"},
}



$ go run main.go
map[a:{Peter} b:{Seth}]
Add
Now, let's see how we can add a value to our map.



func main() {
	var m = map[string]User{
		"a": {"Peter"},
		"b": {"Seth"},
	}
	m["c"] = User{"Steve"}
	fmt.Println(m)
}



$ go run main.go
map[a:{Peter} b:{Seth} c:{Steve}]
Retrieve
We can also retrieve our values from the map using the key.



...
c := m["c"]
fmt.Println("Key c:", c)



$ go run main.go
key c: {Steve}
What if we use a key that is not present in the map?



...
d := m["d"]
fmt.Println("Key d:", d)
We will get the zero value of the map's value type.



$ go run main.go
Key c: {Steve}
Key d: {}
Exists
When you retrieve the value assigned to a given key, it returns an additional boolean value as well. The boolean variable will be true if the key exists, and false otherwise.
Let's try this in an example:



...
c, ok := m["c"]
fmt.Println("Key c:", c, ok)
d, ok := m["d"]
fmt.Println("Key d:", d, ok)



$ go run main.go
Key c: {Steve} Present: true
Key d: {} Present: false
Update
We can also update the value for a key by simply re-assigning a key.



...
m["a"] = "Roger"



$ go run main.go
map[a:{Roger} b:{Seth} c:{Steve}]
Delete
Or, we can delete the key using the built-in delete function.
Here's how the syntax looks:



...
delete(m, "a")
The first argument is the map, and the second is the key we want to delete.
The delete() function doesn't return any value. Also, it doesn't do anything if the key doesn't exist in the map.



$ go run main.go
map[a:{Roger} c:{Steve}]
Iteration
We can iterate over maps with the range keyword like arrays or slices.
Example 1



for key, value := range m {
    fmt.Println(key, value)
}
Example 2



package main
import "fmt"
func main() {
	var m = map[string]User{
		"a": {"Peter"},
		"b": {"Seth"},
	}
	m["c"] = User{"Steve"}
	for key, value := range m {
		fmt.Println("Key: %s, Value: %v", key, value)
	}
}



$ go run main.go
Key: c, Value: {Steve}
Key: a, Value: {Peter}
Key: b, Value: {Seth}
Note that a map is an unordered collection, and therefore the iteration order of a map is not guaranteed to be the same every time we iterate over it.
Properties
Maps are reference types, meaning when we assign a map to a new variable, they refer to the same underlying data structure.
Therefore, changes done by one variable will be visible to the other.



package main
import "fmt"
type User struct {
	Name string
}
func main() {
	var m1 = map[string]User{
		"a": {"Peter"},
		"b": {"Seth"},
	}
	m2 := m1
	m2["c"] = User{"Steve"}
	fmt.Println(m1) // Output: map[a:{Peter} b:{Seth} c:{Steve}]
	fmt.Println(m2) // Output: map[a:{Peter} b:{Seth} c:{Steve}]
}
