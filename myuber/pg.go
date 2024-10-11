package main

import "fmt"

func addOne(nums *[]int) {
	// point is, nums here gets the updated array, not the original empty array
	*nums = append(*nums, 1)
	fmt.Printf("nums inside addone %v\n", *nums)
}

func addTwo(nums *[]int) {
	// defer a funtion call and check the contents of the argument
	defer addOne(nums)
	*nums = append(*nums, 2)
}
func main() {
	nums := make([]int, 0, 2)
	addTwo(&nums)
	fmt.Println(nums) // [2, 1]
}
