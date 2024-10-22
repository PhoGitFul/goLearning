EXERCISE ANSWERS

Fill in the blanks : “A slice is a pointer to an underlying array”.

How to find the index of an element in a slice? sliceName[index] will give the value present at that particular index if its valid index.
However to find the index of a particular element in a slice would need to loop through each element and then whichever slice matches the value being searched for then the index of that element will then be known.  Could use the len function to find out how many elements are in the slice first.

Internally, what is behind a slice? An underlying array, the slice is just a pointer to the array, and may not be pointing to the whole array either could just be a "slice" of it.

What is the capacity of names in the following example : names := []string{"John", "Bob", "Claire", "Nik"}. Explain why. 
Currently capacity is 4 but names can be appended to add more "names". If a new name were appended to the current slice, the capacity of names would increase to 8 and the length to 5.  As a whole new slice of length 4 would be added to the cpacity, increasing it to 8.

What is the capacity of s2 in the following example : s2 := names[0:2]. Explain why. As the slice is starting at array position 0 then the slice will take the capacity of the underlying names array which is 4.  If starting position were 1 (e.g. s2 := names[1:2]) then the capacity of the slice would be array capacity-1 i.e. 3.
The length will be 2 as values in indexes 0 & 1 will be copied to the slice. 

What happens internally when you append an element to a slice with a capacity equals to its length? 
Another slice of the same length will be appended to the slice. So capacity will now equal 2*original length. Length will be increased by how many elements more you appended, so if cpacity and length were both 10 and you appended another 2 elements cpacity would increase to 20 and length becomes 12.

CODE COMMENTS

In the code exercise I show the long way to declare an array, by declaring the day integer array.

I use range in the for loop to print all elements of the array.
	for i, day := range days {
		fmt.Println(i, day)
	}

I make use of the len command to calculate and then print the length of the array.
	fmt.Println("Length of days array is:", len(days))

I then access one element of the array and print that out.
	fmt.Println("Middle of the array is position 3 which is actually day", days[3])

I amend one of the elements of the array and use diff print method to display the array.
	days[5] = 7
	fmt.Println(days)

Finally, as part of the formal exercise, I create a slice of integers called fibs, print it out then append some more values to the slice and print again.
	fibs := []int{0, 1, 1, 2, 3, 5}
	fmt.Println("\nFibs")
	fmt.Println(fibs)
	fibs = append(fibs, 8, 13, 21)
	fmt.Println(fibs)

Extra bits I performed in the exercise for my learning benefit. 

Short hand way of declaring an array, then print out specific elements of that array
	weekday := [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println("\n", weekday[0], ", day", days[0], ", is the first day of the week and so", weekday[3], "is the middle of the week and is day", days[3])

Declare and set an array of structs
	type dayOfWeek struct {
		day     int
		weekday string
	}
	daysOfWeek := [7]dayOfWeek{
		{1, "Sun"},
		{2, "Mon"},
		{3, "Tue"},
		{4, "Wed"},
		{5, "Thu"},
		{6, "Fri"},
		{7, "Sat"},
	}

Print out the struct array
	for i, dayOfWeek := range daysOfWeek {
		fmt.Println(i, dayOfWeek)
	}
	fmt.Println("\nI wish the weekends were longer!")

Amend some elements of the array and print out for range loop.
	daysOfWeek[5] = dayOfWeek{day: 6, weekday: "Sat"}
	daysOfWeek[0] = dayOfWeek{day: 1, weekday: "Sat"}
	daysOfWeek[1] = dayOfWeek{day: 2, weekday: "Sun"}

	for i, dayOfWeek := range daysOfWeek {
		fmt.Println(i, dayOfWeek)
	}

Create a slice of this array and print out the slice.
	daysSlice := weekday[5:] // creates a slice from index 6 to 7
	fmt.Println("\nFriday and Saturday are my favourite days.", daysSlice)

I also update an element in ths slice to show that updating a slice based on an array actually updates the underlying array data.
The slice is just a pointer to the array.