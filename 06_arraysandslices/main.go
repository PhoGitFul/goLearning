package main

import "fmt"

func main() {

	var scores [5]int
	scores[0] = 90
	scores[1] = 80
	scores[2] = 70
	scores[3] = 60
	scores[4] = 50

	fruits := [3]string{"apple", "banana", "cherry"}

	type Person struct {
		name string
		age  int
	}
	people := [2]Person{
		{"John", 25},
		{"Jane", 30},
	}
	fmt.Println(scores[0])
	fmt.Println(fruits[1])
	fmt.Println(people[1].name)
	fmt.Println()

	s1 := scores[1:4]
	fmt.Println("slice s1:", s1)
	fmt.Println("element 1:", s1[1])
	fmt.Println("slice s1 length is", len(s1))
	fmt.Println("slice s1 cap is", cap(s1))
	fmt.Println()

	s1 = append(s1, 40)
	fmt.Println("After appending a value of 40 slice s1 is now:", s1)
	fmt.Println("s1 length is", len(s1))
	fmt.Println("s1 cap is", cap(s1))

	s1 = append(s1, 3, 2)
	fmt.Println("After further appending two more elements 3 & 2 s1 is now", s1)
	fmt.Println("s1 length is", len(s1))
	fmt.Println("s1 cap is", cap(s1))
	fmt.Println()
	fmt.Println("s1 slice is", s1[0], s1[1], s1[2], s1[3], s1[4], s1[5])
	fmt.Println()

	morescores := []int{9, 8, 7, 6, 5}
	fmt.Println(morescores)
	fmt.Println("morescores element 2 is:", morescores[2])
	fmt.Println("morescores length is:", len(morescores))
	fmt.Println("morescores cap is:", cap(morescores))
	fmt.Println()

	morescores = append(morescores, 4)
	fmt.Println(morescores)
	fmt.Println("This first append increases morescores length by 1 to:", len(morescores))
	fmt.Println("First append increases morescores cap by 2 as another slice of same length added, cap:", cap(morescores))
	fmt.Println()

	s2 := scores[0:1]
	fmt.Println("slice s2[0:1]:", s2)
	fmt.Println("slice s2 length is", len(s2))
	fmt.Println("slice s2 cap is", cap(s2))
	fmt.Println()

	s3 := scores[0:2]
	fmt.Println("slice s3[0:2]:", s3)
	fmt.Println("slice s3 length is", len(s3))
	fmt.Println("slice s3 cap is", cap(s3))
	fmt.Println()

	s4 := scores[0:4]
	fmt.Println("slice s4[0:4]:", s4)
	fmt.Println("slice s4 length is", len(s4))
	fmt.Println("slice s4 cap is", cap(s4))
	fmt.Println()

	s5 := scores[1:4]
	fmt.Println("slice s5[1:4]:", s5)
	fmt.Println("slice s5 length is", len(s5))
	fmt.Println("slice s5 cap is", cap(s5))
	fmt.Println()

	s6 := scores[2:4]
	fmt.Println("slice s6[2:4]:", s6)
	fmt.Println("slice s6 length is", len(s6))
	fmt.Println("slice s6 cap is", cap(s6))
	fmt.Println()

	s7 := scores[3:4]
	fmt.Println("slice s7[3:4]:", s7)
	fmt.Println("slice s7 length is", len(s7))
	fmt.Println("slice s7 cap is", cap(s7))
	fmt.Println()

}
