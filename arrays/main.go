package main

import "fmt"

func main() {
	// declaring and initializing an array
	var nums [3]int
	nums[0] = 10
	nums[1] = 20
	nums[2] = 30
	fmt.Println("Array:", nums)

	// declaring and initializing an array with shorthand
	var numbers [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array:", numbers)
}
