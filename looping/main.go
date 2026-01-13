package main

import "fmt"
func main() {
	//while loop using tradition for loop
	i := 5;
	for i < 10 {
		fmt.Println("Iteration:", i)
		i++
	}

	// infinite loop
	// for {
	// 	fmt.Println("Infinite Loop")
	// }

	// traditional for loop
	for j := 1; j <= 5; j++ {
		fmt.Println("Iteration:", j)
	}
	// using range 
	for k := range 3 {
		fmt.Println("Iteration:", k)
	}
	// using break and continue
	for l := 1; l <= 10; l++ {
		if l%2 == 0 {
			continue // skip even numbers
		}
		if l > 7 {
			break // exit loop if l is greater than 7
		}
		fmt.Println("Odd number less than 8:", l)
	}



	// for i := 1; i <= 5; i++ {
	// 	fmt.Println("Iteration:", i)
	// }
}