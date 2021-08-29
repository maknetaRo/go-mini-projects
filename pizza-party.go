package main

import (
	"fmt"
	"strconv"
	"./calculator"
)

func main() {
	var people, pizzas string
	piecesPerPizza := 8
	fmt.Println("How many people?")
	fmt.Scanln(&people)
	peopleInt, err := strconv.Atoi(people) 
	for err != nil {
		fmt.Println("Please write an integer")
		fmt.Scanln(&people)
		peopleInt, err = strconv.Atoi(people)
	}


	fmt.Println("How many pizzas do you have?")
	fmt.Scanln(&pizzas)
	pizzasInt, err := strconv.Atoi(pizzas)
	for err != nil {
		fmt.Println("Please write an integer")
		fmt.Scanln(&pizzas)
		pizzasInt, err = strconv.Atoi(pizzas)
	}

	fmt.Printf("%v people with %v pizzas.\n", people, pizzas)
	pieces := calculator.DivideInt(pizzasInt * piecesPerPizza, peopleInt)
	fmt.Printf("Each person gets %v pieces of pizza.\n", pieces)
	leftOvers := calculator.Reminder(pizzasInt * piecesPerPizza, peopleInt)
	fmt.Printf("There are %v leftover pieces.\n", leftOvers)
}