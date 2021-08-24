package main

import "fmt"

func main() {
	fmt.Println("What's your name?")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello, " + name + ", nice to meet you!")
}