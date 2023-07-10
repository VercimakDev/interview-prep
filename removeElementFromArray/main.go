package main

import (
	"fmt"
)

func main() {
	var nums = []int{0,1,2,2,3,0,4,2}
	val := 2
	removeElement(nums, val)
}

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++{
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	fmt.Print(nums)
	return len(nums)
}
