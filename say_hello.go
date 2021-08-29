package main

import "fmt"

func SayHello() string {
	fmt.Println("What's your name?")
	var name string
	fmt.Scanln(&name)
	return "Hello, " + name + ", nice to meet you!"
}

func main() {
	fmt.Println(SayHello())
}