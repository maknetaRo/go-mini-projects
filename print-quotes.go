package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func Scanf(a *string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine() 
	*a = string(data)
}

func main() {
	var quote string 
	var author string 
	
	fmt.Println("What is the quote?")
	Scanf(&quote)
	
	fmt.Println("Who said it?")
	Scanf(&author)

	
	fmt.Printf("%s says, \"%s\".\n", strings.Title(author), quote)	

}