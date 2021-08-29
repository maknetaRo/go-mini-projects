package main

import (
	"./calculator"
	"fmt"
	"strings"

)

func main() {
	var length, width float64
	var feetOrMeters string

	fmt.Println("Feet? or Meters?")
	fmt.Scanln(&feetOrMeters)

	if strings.ToLower(feetOrMeters) == "feet" {
		fmt.Println("What is the length of the room in feet?")
		fmt.Scanln(&length)

		fmt.Println("What is the width of the room in feet?")
		fmt.Scanln(&width)

		feetArea := calculator.Multiply(length, width)

		meterSquare := calculator.FeetToMeters(feetArea)
		fmt.Printf("You entered dimensions of %v feet by %v feet.\n", length, width)
		fmt.Printf("The area is: \n %v square feet \n %v square meters.\n", feetArea, meterSquare)
	} else {
		fmt.Println("What is the length of the room in meters?")
		fmt.Scanln(&length)

		fmt.Println("What is the width of the room in meters?")
		fmt.Scanln(&width)

		meters := calculator.Multiply(length, width)

		feet := calculator.MetersToFeet(meters)
		fmt.Printf("You entered dimensions of %v meters by %v meters.\n", length, width)
		fmt.Printf("The area is: \n %v square meters \n %v square feet.\n", meters, feet)
	}

}
