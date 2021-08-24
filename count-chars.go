package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("What is the input string? ")
	var word string 
	fmt.Scanln(&word)
	for len(word) == 0 {
		fmt.Println("You need to write a word.")
		fmt.Scanln(&word)		
	}
	word = strings.Title(word)
	var word_length int 
	word_length = len(word)
	response := fmt.Sprintf("%s has %d characters.", word, word_length )
	fmt.Println(response)	
	
}