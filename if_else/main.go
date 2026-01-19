package main

import "fmt"

func main() { // main function is the entry point of the program
	fmt.Println("Hello, World!") // Print "Hello, World!" to the console
	if age := 20; age >= 18 {
		fmt.Println("You are an adult.") // Print if age is 18 or older
	} else {
		fmt.Println("You are a minor.") // Print if age is under 18
	}
	//we dont have the ternary operator in go
	//
}