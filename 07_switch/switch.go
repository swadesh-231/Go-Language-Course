package main

import (
	"fmt"
)

func main() {
	// i := 1

	// switch i {
	// case 1:
	// 	fmt.Println("i is 1")
	// default:
	// 	fmt.Println("no match")
	// }

	//multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("its weekend")
	// default:
	// 	fmt.Println("")
	// }

	// switch day := "Sat"; day {
	// case "Sat", "Sun":
	// 	fmt.Println("Weekend")
	// default:
	// 	fmt.Println("Weekday")
	// }


	//type switch 

	whoAmI := func(i interface{}){
		switch t := i.(type){
		case int:
			fmt.Println("Integer")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("boolean")
		default:
			fmt.Println("other",t)
		}
	}
		whoAmI(false)

}
