package main

import "fmt"

func main() {
	// Static Type declarations
	//var barName string
	//var barType, barLocation string
	//var barRating int

	//barName = "Welcome to 46 Ounces!"
	//barType = "Spanish"
	//barLocation = "Bangalore"
	//barRating = 5

	//var barName string = "46 Ounces"
	//var barType = "German"
	//var barLocation = "Bangalore"
	//var barRating = 4.5

	// Dynamic Type Declarations
	//barName := "46 Ounces!"
	//barType := "Spanish"
	//barLocation := "Bangalore"
	//barRating := 4.5

	// fmt.Println("Welcome to 46 Ounces")

	// fmt.Println(barName + barType + barLocation)
	// fmt.Println(barName+" "+barType+" "+barLocation, barRating)
	// fmt.Println(barName, barType, barLocation, barRating)

	// Mixed Variable Declaration
	var barName, barType, barLocation, barRating = "46 Ounces", "German", "Bangalore", 4.5

	fmt.Println(barName, barType, barLocation, barRating)
	fmt.Printf("a is of type %T\n", barRating)

	const barOwner = "Malhotra"
	fmt.Println(barOwner)
	//barOwner = "Prabhat"
	//fmt.Println(barOwner)

}
