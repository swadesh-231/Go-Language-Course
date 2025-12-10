package main

import "fmt"

//slices -> dynamic array
func main(){

	// un initialized slice is nill
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var number = make([]int,1,5)
	// fmt.Println(len(number))
	// fmt.Println(cap(number))
	// number[0] = 23
	// number = append(number, 1,2,3,4,5,6,7,8)
	// fmt.Println(number)


	// copy function
	// Create an empty int slice with length 0 and capacity 5
	var nums = make([]int,0,5)
	nums = append(nums, 2)

	var nums2 = make([]int ,len(nums))
	

	//copy function 
	copy(nums2,nums)
	fmt.Println(nums,nums2)

	//slice operator

}