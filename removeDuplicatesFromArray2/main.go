package main

import "fmt"

func main() {
	var nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	k := removeDuplicates(nums)
	fmt.Println(nums, k)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	i := 0
	double := false
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
			double = false
		} else {
			if !double {
				double = true
				i++
				nums[i] = nums[j]
			}
		}
	}
	return i + 1
}
