package main

import "fmt"

func main() {
	var nums = []int{1,2,3}
	k := removeDuplicates(nums)
	fmt.Println(nums, k)
}

func removeDuplicates(nums []int) int {
	var queue []int
	last_int := nums[0]
	unique := 1
	for i := 1; i < len(nums); i++ {
		if last_int == nums[i] {
			queue = append(queue, i)
		} else {
			if len(queue) != 0 {
				nums[queue[0]] = nums[i]
				queue = queue[1:]
				queue = append(queue, i)
			}
			last_int = nums[i]
			unique++
		}
	}
	return unique
}
