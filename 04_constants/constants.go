package main

import "fmt"

const city = "Kolkata"

// age:=23 // this things not allowed outside function

func main() {
	const name = "GoLang" // for constants in go
	fmt.Println(name)
	fmt.Println(city)

	const (
		port = 3030
		host = "local"
	)

	fmt.Println(port,host)

}
