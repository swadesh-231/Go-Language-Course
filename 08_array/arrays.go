package main

import "fmt"

func main() {
	// int-> 0 bool->false string->""
	var nums [4]int
	nums[0] = 1

	fmt.Println(len(nums))
	fmt.Println(nums[0])
	fmt.Println(nums)

	var flags [4]bool
	fmt.Println(flags)

	var city [3]string
	city[0] = "kolkata"
	fmt.Println(city)

	numbers := [4]int{1, 2, 3, 4} // to declare it in single lien
	fmt.Println(numbers)

	//2d array in go

	var arr [3][4]int

	fmt.Println(arr)

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix)

}
