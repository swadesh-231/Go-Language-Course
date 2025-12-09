package main

import "fmt"

// for only construct in go for looping
func main() {
	// while style loops
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
	fmt.Println()
	//classic for loops 
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

}
