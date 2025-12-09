package main

import (
	"fmt"
)

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
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println()
	//range in go language 
	//0 -> 4 
	for i := range 5 {
		fmt.Println(i)
	}
}
