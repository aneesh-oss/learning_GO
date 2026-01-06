package main

import "fmt" // import statement is used to include other packages

const pi float64 = 3.14159
var radius int = 7

const (
	host string = "localhost"
	port int = 8080
)


func main() { // main function is the entry point of the program
	var message string = "Hello, Variables!"
	fmt.Println(message) // Print the message variable to the console

	var num1 int = 15
	var num2 int = 25
	var sum int = num1 + num2
	fmt.Println("Sum:", sum) // Print the sum of num1 and num2
	area := pi * float64(radius*radius)
	fmt.Println("Area of circle with radius", radius, "is:", area) // Print the area of the circle
	fmt.Println("Server running at", host, "on port", port) // Print server details

}