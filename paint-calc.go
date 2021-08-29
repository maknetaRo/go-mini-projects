package main

import (
	"fmt"

	"./calculator"
)



func comparison(a int, b float64) bool {
	aFloat := float64(a)
	return aFloat < b
}

func main() {
	var gallons float64
	var coverage float64 = 360
	var width, length float64
	fmt.Println("What is the width of the room in feet?")
	fmt.Scanln(&width)

	fmt.Println("What is the length of the room in feet?")
	fmt.Scanln(&length)

	feetSquare := calculator.Multiply(width, length)
	gallons = calculator.Divide(feetSquare, coverage)
	gallonsInt := int(gallons)

	if comparison(gallonsInt, gallons) {
		gallonsInt = gallonsInt + 1
	}


	fmt.Printf("You will need to purchase %v gallons of paint to cover %v square feet.\n", gallonsInt, feetSquare)
}