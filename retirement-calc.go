package main

import (
	"fmt"
	"strconv"
	"time"
) 

func main() {
	var currentAge, retirementAge string
	var yearsLeft, retirementYear int 

	dt := time.Now()
	year := dt.Year()	
	
	fmt.Println("What is you current age?")
	fmt.Scanln(&currentAge)
	currentAgeInt, err := strconv.Atoi(currentAge)
	for err != nil {
		fmt.Println("You should have given an integer.")
		fmt.Println("What is you current age?")
		fmt.Scanln(&currentAge)
		currentAgeInt, err = strconv.Atoi(currentAge)
	}

	fmt.Println("At what age would you like to retire?")
	fmt.Scanln(&retirementAge)
	retirementAgeInt, err := strconv.Atoi(retirementAge)
	for err != nil {
		fmt.Println("You should have given an integer.")
		fmt.Println("At what age would you like to retire?")
		fmt.Scanln(&retirementAge)
		retirementAgeInt, err = strconv.Atoi(retirementAge)
	}


	yearsLeft = retirementAgeInt - currentAgeInt
	retirementYear = year + yearsLeft

	if yearsLeft < 0 {
		fmt.Println("You should have retired " + strconv.Itoa(yearsLeft)[1:] + " years ago.")
		fmt.Println("It's " + strconv.Itoa(year) + ", so you should have retired in " + strconv.Itoa(retirementYear) + ".")
	} else {
		fmt.Println("You have " + strconv.Itoa(yearsLeft) + " years left until you can retire.")
		fmt.Println("It's " + strconv.Itoa(year) + ", so you can retire in " + strconv.Itoa(retirementYear) + ".")
	
	}

	
	// fmt.Println("***********************")
	// fmt.Println("Another way to convert integers to strings with fmt.Sprintf function")
	// fmt.Println("***********************")
	// message1 := fmt.Sprintf("You have %d years left until you can retire.", yearsLeft)
	// message2 := fmt.Sprintf("It's %d, so you can retire in %d.", year, retirementYear)
	// fmt.Println(message1)
	// fmt.Println(message2)
}