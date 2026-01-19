package main

import "fmt"

func main() { // main function is the entry point of the program
	fmt.Println("Hello, World!") // Print "Hello, World!" to the console
	//switch statement
	//we use switch when we have multiple conditions to check
	//and we want to execute different code based on which condition is met
	//it is more readable than using multiple if-else statements
	//we dont use the break statement in go switch cases
	//because go automatically breaks out of the switch after executing a case
	//unless we use the fallthrough statement to explicitly continue to the next case
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}
	//``switch with multiple values in a case
	letter := "a"
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("Vowel")
	default:
		fmt.Println("Consonant")
	}
	//switch without an expression
	num := 15
	switch {
	case num < 10:
		fmt.Println("Less than 10")
	case num >= 10 && num < 20:
		fmt.Println("Between 10 and 20")
	default:
		fmt.Println("20 or more")
	}
	//switch with fallthrough
	n := 2
	switch n {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
	}
	//what happens in the fallthrough is that once case 2 is matched,
	//it prints "Two" and then continues to the next case (case 3)
	//and prints "Three" as well, regardless of whether case 3 matches or not

	//what is the use of switch statements in go?
	//switch statements are used to simplify complex conditional logic
	//they improve code readability and maintainability
	//they allow for cleaner handling of multiple conditions compared to if-else chains

	//what happens if no case matches and there is no default case?
	//if no case matches and there is no default case,
	//the switch statement simply does nothing and control moves to the next statement after the switch

	//type switch
	//what is interface{} in go?
	//interface{} is the empty interface in go
	//it can hold values of any type
	//every type in go implements at least zero methods, so they all satisfy the empty interface
	//this makes interface{} a powerful tool for writing generic code that can work with any type
	var x interface{} = 42
	switch x.(type) {
	case int:
		fmt.Println("x is an integer")
	case string:
		fmt.Println("x is a string")
	default:
		fmt.Println("x is of a different type")
	}
	//explian type switch
	//a type switch is used to compare the dynamic type of an interface variable
	//it allows us to execute different code based on the actual type of the variable at runtime
	//this is useful when working with interfaces and we need to handle different types differently
	//in the example above, we check if x is of type int or string and print accordingly
}