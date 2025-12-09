package main

import "fmt"

func main(){
	var name string = "Hello World"
	var age =23
	fmt.Println((name))
	fmt.Println(age)

	 // Explicit type
    var a int = 10

    // Type inference
    var b = 20

    // Short declaration (ONLY inside functions)
    c := 30
	  // Multiple variables
    var x, y int = 40, 50
    username, life := "John", 25

    fmt.Println(a, b, c, x, y, username, life)

	var city string
	city = "Kolkata"
	fmt.Println(city)
}