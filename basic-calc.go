package main

import (
	"calculator"
	"fmt"
)

func Print_operations() {
	fmt.Println("Select operation.")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")	
 }


func main() {
	var operation int
	Print_operations()

	fmt.Scanln(&operation)
	var first_num, second_num float64
	
	fmt.Println("Enter 1st Number: ")
	fmt.Scanln(&first_num)
	fmt.Println("Enter 2nd Number: ")
	fmt.Scanln(&second_num)

	switch operation {
		case 1:
			fmt.Println(calculator.Add(first_num, second_num))
		case 2:
			fmt.Println(calculator.Subtract(first_num, second_num))
		case 3:
			fmt.Println(calculator.Multiply(first_num, second_num))
		case 4:
			fmt.Println(calculator.Divide(first_num, second_num))
		default:
			fmt.Println("Error. I didn't know what operations you wanted me to perform."); 			
	}


}