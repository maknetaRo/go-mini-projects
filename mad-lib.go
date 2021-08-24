package main

import "fmt"

func main() {
	var noun, verb, adjective, adverb string
	fmt.Println("Enter a noun: ")
	fmt.Scanln(&noun)
	fmt.Println("Enter a verb: ")
	fmt.Scanln(&verb)
	fmt.Println("Enter an adjective: ")
	fmt.Scanln(&adjective)
	fmt.Println("Enter an adverb: ")
	fmt.Scanln(&adverb)

	fmt.Println("Do you " + verb + " your " + adjective + " " + noun + " " + adverb +"? That's hilarious!")
}