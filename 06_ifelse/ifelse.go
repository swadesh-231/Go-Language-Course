package main

import (
	"fmt"
)

func main() {
	// age := 18

	// if age > 18 {
	// 	fmt.Println("Adult")
	// } else if age == 18 {
	// 	fmt.Println("Just Adult")
	// } else {
	// 	fmt.Println("Minor")
	// }

	// if num := 10; num%2 == 0 {
	// 	fmt.Println("Even")
	// } else {
	// 	fmt.Println("Odd")
	// }

	// var role = "admin"
	// var hasPermission = false

	// if role == "admin" && hasPermission  {
	// 	fmt.Println("Welcome to My Page")
	// }

	if age := 15; age >= 18 {
		fmt.Println("Person is adult", age)
	} else if age >= 13 {
		fmt.Println("Teenager", age)
	} else {
		fmt.Println("Child", age)
	}

	// go don't have ternary operators

}
