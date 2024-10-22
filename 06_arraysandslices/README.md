CODE COMMENTS

The main program show examples and uses of arrays and slices.

You can declare an array like this:

	var scores [5]int
	scores[0] = 90
	scores[1] = 80
	scores[2] = 70
	scores[3] = 60
	scores[4] = 50

    or Shorthand declaration like this:
    fruits := [3]string{"apple", "banana", "cherry"}

Arrays can hold Structs:
	type Person struct {
		name string
		age  int
	}
	people := [2]Person{ //people is an array of Person which is a Struct
		{"John", 25},
		{"Jane", 30},
    }

Examples of referencing the above array elements
    fmt.Println(scores[0])      // prints 90
	fmt.Println(fruits[1])      // prints "banana"
	fmt.Println(people[1].name) // prints "Jane"

Declare a slice so no number is required in the brackets []
	morescores := []int{90, 80, 70, 60, 50}
	fmt.Println(morescores[2])  // prints 70
a slice usually references an underlying array and will slice part of it:
s1 is a slice of the array scores taking just 3 elements with indexes from 1 to 3
	s1 := scores[1:4]
    Whilst the slice takes the values from indexes 1, 2 & 3 so it has length of 3, the capacity is 4 as the starting position is 1, if starting pos was 0 the cap would be 5, same as the underlying array capacity.

a slice can be appended with new values, with the append command

	s1 = append(s1, 40)
As s1 was not yet at capacity as we sliced it from the array but started at element 1 then we have 1 free element in which to fill up so capacity will be 5 once we ahve appended this new value.t

When we append 2 new values as we were at capacity then the new capacity will double the original capacity, as a new slice of cap 4 will be added so we now have cap 8, and the length will increase by 2 to 6.

To show that the capacity of a slice depends on the position in the array the slice starts, I put through a number of examples of slices to illustrate that the capacity of the slice is taken from the starting point of the slice, so if from 0, the cpacity will match the capacity of the array, if from 1 it will be array capacity-1 and so on.