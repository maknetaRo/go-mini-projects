package main

import (
	"./calculator"
	"fmt"
)

const firstPrompt = "Enter the principal: "
const secondPrompt = "Enter the rate of interest: "
const thirdPrompt = "Enter the number of years: "

func main() {
	var principal, rate, years float64
	fmt.Println(firstPrompt)
	fmt.Scanln(&principal)
	fmt.Println(secondPrompt)
	fmt.Scanln(&rate)
	fmt.Println(thirdPrompt)
	fmt.Scanln(&years)
	interest := calculator.CalcInterest(principal, rate, years)

	fmt.Printf("After %v years at %v%%, the investment will be worth $%v.\n", years, rate, interest)

}
