package main // This needs to be set to main
//Otherwise the exe will not run!

import "fmt" // We are importing the fmt code base and can use the fns in our code

func main() {
	var stationName string //set the variable types here
	var nextTrainTime int
	var isUptownTrain bool

	stationName = "Union Square"
	nextTrainTime = 12
	isUptownTrain = false //from looking at output I know this needs to initially be set to false

	// This next section will use the Println function from fmt "package"?
	fmt.Println("Current station:", stationName) //add variable
	fmt.Println("Next train:", nextTrainTime, "minutes")
	fmt.Println("Is uptown:", isUptownTrain)

	//below set the variables for the next section
	stationName = "Grand Central"
	nextTrainTime = 3
	isUptownTrain = true

	fmt.Println("")
	fmt.Println("Current station:", stationName)
	fmt.Println("Next train:", nextTrainTime, "minutes")
	fmt.Println("Is uptown:", isUptownTrain)
}
