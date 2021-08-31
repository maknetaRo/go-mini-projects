package main 

import "fmt"

const firstPrompt = "Enter the principal: "
func Prompt(prompt string)  {
	var result int
	fmt.Println(prompt)
	fmt.Scanln(result)
}

func main() {
	Prompt(firstPrompt)
}