package main

import "fmt"

func main() {

	var days [7]int
	days[0] = 1
	days[1] = 2
	days[2] = 3
	days[3] = 4
	days[4] = 5
	days[5] = 6
	days[6] = 7

	for i, day := range days {
		fmt.Println(i, day)
	}
	fmt.Println("Length of days array is:", len(days))
	fmt.Println("Middle of the array is position 3 which is actually day", days[3])

	fmt.Println("If only we could have Saturday the 7th day earlier")
	days[5] = 7
	fmt.Println(days)

	fibs := []int{0, 1, 1, 2, 3, 5}
	fmt.Println("\nFibs")
	fmt.Println(fibs)
	fibs = append(fibs, 8, 13, 21)
	fmt.Println(fibs)

	//Performed extra examples of arrays and slices
	weekday := [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println("\n", weekday[0], ", day", days[0], ", is the first day of the week and so", weekday[3], "is the middle of the week and is day", days[3])

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

	for i, dayOfWeek := range daysOfWeek {
		fmt.Println(i, dayOfWeek)
	}
	fmt.Println("\nI wish the weekends were longer!")

	daysOfWeek[5] = dayOfWeek{day: 6, weekday: "Sat"}
	daysOfWeek[0] = dayOfWeek{day: 1, weekday: "Sat"}
	daysOfWeek[1] = dayOfWeek{day: 2, weekday: "Sun"}

	for i, dayOfWeek := range daysOfWeek {
		fmt.Println(i, dayOfWeek)
	}

	daysSlice := weekday[5:] // creates a slice from index 6 to 7
	fmt.Println("\nFriday and Saturday are my favourite days.", daysSlice)
	fmt.Println()

	daysSlice[0] = "Funday"
	fmt.Println("Update daysSlice[0] to Funday, so daysSlice is now:", daysSlice)
	fmt.Println()
	fmt.Println("As daysSlice[] is just referencing the underlying array weekday[], daysSlice[0] points to weekday[5]")
	fmt.Println()
	fmt.Println("So updating daysSlice is actually updating weekday array.\n Which now looks like this:", weekday)

}
